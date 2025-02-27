package main

import (
	"log"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/config"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/handlers"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/routes"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/services"
)

func main() {
	log.Println("Setting up Database")
	db := config.InitDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	gigRepo := repositories.NewGigRepository(db)
	gigService := services.NewGigService(gigRepo)
	gigHandler := handlers.NewGigHandler(gigService)

	r := routes.SetupRouter(userHandler, gigHandler)

	r.Run("0.0.0.0:8080") // Start server on port 8080
}
