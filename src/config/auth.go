package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func InitCognitoClient() *cognitoidentityprovider.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}
	return cognitoidentityprovider.NewFromConfig(cfg)
}

func GetUserPoolID() string {
	return os.Getenv("COGNITO_USER_POOL_ID")
}

func GetClientID() string {
	return os.Getenv("COGNITO_CLIENT_ID")
}

func GetRegion() string {
	return os.Getenv("AWS_REGION")
}
