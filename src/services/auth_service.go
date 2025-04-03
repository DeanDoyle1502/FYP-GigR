package services

import (
	"context"
	"fmt"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/config"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

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

	// ✅ Retrieve Cognito sub
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

	// ✅ Store full user info
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

	// Return the JWT ID token
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
