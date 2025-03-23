package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/config"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/handlers"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/middleware"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/routes"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	middleware.SetupJWKs()

	log.Println("Setting up Database")
	db := config.InitDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	cognitoClient := config.InitCognitoClient()
	authService := services.NewAuthService(cognitoClient, userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	gigRepo := repositories.NewGigRepository(db)
	gigService := services.NewGigService(gigRepo, authService)
	gigHandler := handlers.NewGigHandler(gigService)

	r := routes.SetupRouter(userHandler, gigHandler, authHandler)

	fmt.Println("🚀 Server started with auth routes")

	r.Run("0.0.0.0:8080") // Start server on port 8080
}
