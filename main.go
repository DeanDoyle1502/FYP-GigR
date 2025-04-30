package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/config"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/handlers"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/routes"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// ✅ Initialize Cognito JWKs for token verification
	if err := services.SetupJWKs(); err != nil {
		log.Fatalf("Failed to set up Cognito JWKs: %v", err)
	}
	log.Println("✅ Cognito JWKs loaded")

	// Database setup
	log.Println("Setting up Database")
	db := config.InitDB()

	log.Println("Setting up DynamoDB")
	dynamoClient := config.InitDynamoDB()
	config.EnsureDynamoTablesExist()

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	gigRepo := repositories.NewGigRepository(db)
	messageRepo := repositories.NewMessageRepository(dynamoClient)
	chatSessionRepo := repositories.NewChatSessionRepository(dynamoClient, "gigrChatSessions")

	// Services
	userService := services.NewUserService(userRepo)
	cognitoClient := config.InitCognitoClient()
	authService := services.NewAuthService(cognitoClient, userRepo)
	gigService := services.NewGigService(gigRepo, authService)
	messageService := services.NewMessageService(messageRepo)
	chatSessionService := services.NewChatSessionService(chatSessionRepo, gigRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	gigHandler := handlers.NewGigHandler(gigService)
	messageHandler := handlers.NewMessageHandler(messageService)
	chatSessionHandler := handlers.NewChatSessionHandler(chatSessionService)

	r := gin.Default()

	// CORS setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://gigr-frontend.s3-website-eu-west-1.amazonaws.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	routes.SetupRouter(r, userHandler, gigHandler, authHandler, messageHandler, chatSessionHandler, userRepo)

	fmt.Println("🚀 Server started with auth routes")
	r.Run("0.0.0.0:8080") // Start server on port 8080
}
