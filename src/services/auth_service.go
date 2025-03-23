package services

import (
	"context"
	"fmt"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type AuthService struct {
	Cognito *cognitoidentityprovider.Client
}

func NewAuthService(client *cognitoidentityprovider.Client) *AuthService {
	return &AuthService{Cognito: client}
}

func (s *AuthService) RegisterUser(email, password string) error {
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
		return fmt.Errorf("failed to register user: %w", err)
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
