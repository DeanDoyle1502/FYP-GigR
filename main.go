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

	r := routes.SetupRouter(userHandler)

	r.Run(":8080") // Start server on port 8080
}
