package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/config"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var jwks *keyfunc.JWKS

type AuthService struct {
	Cognito  *cognitoidentityprovider.Client
	UserRepo *repositories.UserRepository
}

func NewAuthService(client *cognitoidentityprovider.Client, userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		Cognito:  client,
		UserRepo: userRepo,
	}
}

func (s *AuthService) RegisterUser(email, password, name, instrument, location, bio string) error {
	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(config.GetClientID()),
		Username: aws.String(email),
		Password: aws.String(password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	_, err := s.Cognito.SignUp(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to register user in Cognito: %w", err)
	}

	adminInput := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(config.GetUserPoolID()),
		Username:   aws.String(email),
	}

	userOutput, err := s.Cognito.AdminGetUser(context.TODO(), adminInput)
	if err != nil {
		return fmt.Errorf("user created in Cognito but failed to retrieve details: %w", err)
	}

	var cognitoSub string
	for _, attr := range userOutput.UserAttributes {
		if *attr.Name == "sub" {
			cognitoSub = *attr.Value
			break
		}
	}
	if cognitoSub == "" {
		return fmt.Errorf("could not find Cognito sub after signup")
	}

	user := &models.User{
		Email:      email,
		CognitoSub: cognitoSub,
		Name:       name,
		Instrument: instrument,
		Location:   location,
		Bio:        bio,
	}

	if err := s.UserRepo.CreateUser(user); err != nil {
		return fmt.Errorf("user added to Cognito, but failed to save in DB: %w", err)
	}

	return nil
}

func (s *AuthService) LoginUser(email, password string) (string, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: aws.String(config.GetClientID()),
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
	}

	result, err := s.Cognito.InitiateAuth(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("login failed: %w", err)
	}

	return *result.AuthenticationResult.IdToken, nil
}

func (s *AuthService) ConfirmUser(email, code string) error {
	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(config.GetClientID()),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
	}

	_, err := s.Cognito.ConfirmSignUp(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("confirmation failed: %w", err)
	}

	return nil
}

func (s *AuthService) GetOrCreateUser(sub, email string) (*models.User, error) {
	return s.UserRepo.GetOrCreateByCognitoSub(sub, email)
}

// 🆕 Initialize JWKs on app start
func SetupJWKs() error {
	jwksURL := fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json",
		config.GetRegion(),
		config.GetUserPoolID(),
	)

	var err error
	jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{})
	return err
}

// 🆕 Extract userID from Authorization header
func (s *AuthService) ExtractUserIDFromToken(c *gin.Context) (uint, error) {
	authHeader := c.GetHeader("Authorization")
	fmt.Println("🧪 Authorization header:", authHeader)

	if authHeader == "" {
		fmt.Println("❌ Missing Authorization header")
		return 0, errors.New("missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		fmt.Println("❌ Invalid Authorization header format:", parts)
		return 0, errors.New("invalid Authorization header format")
	}

	tokenString := parts[1]
	fmt.Println("🔐 Token string:", tokenString)

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		fmt.Println("❌ JWT parse failed:", err)
		return 0, errors.New("invalid token")
	}
	if !token.Valid {
		fmt.Println("❌ Token is not valid")
		return 0, errors.New("invalid token")
	}
	fmt.Println("✅ Token parsed successfully")

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("❌ Unable to parse claims")
		return 0, errors.New("invalid claims")
	}
	fmt.Printf("🧾 Claims: %+v\n", claims)

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		fmt.Println("❌ Missing sub in claims")
		return 0, errors.New("missing sub claim")
	}
	fmt.Println("✅ Extracted sub:", sub)

	user, err := s.UserRepo.GetUserByCognitoSub(sub)
	if err != nil {
		fmt.Println("❌ Failed to get user from DB by sub:", err)
		return 0, errors.New("user not found")
	}

	fmt.Println("✅ Found user:", user.ID)
	return user.ID, nil
}
