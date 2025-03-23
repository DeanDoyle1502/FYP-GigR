package routes

import (
	"github.com/DeanDoyle1502/FYP-GigR.git/src/handlers"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, gigHandler *handlers.GigHandler, authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Define user routes
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)

	// Define gig routes
	r.POST("/gigs", gigHandler.CreateGig)
	r.GET("/gigs", gigHandler.GetAllGigs)
	r.GET("/gigs/:id", gigHandler.GetGig)
	r.POST("/gigs/:id/apply", gigHandler.ApplyForGig)
	r.POST("/gigs/:id/accept/:musicianID", gigHandler.AcceptMusicianForGig)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Authorisation
	r.POST("/auth/register", authHandler.RegisterUser)
	r.POST("/auth/login", authHandler.LoginUser)
	r.POST("/auth/confirm", authHandler.ConfirmUser)

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/me", func(c *gin.Context) {
			claims, _ := c.Get("user")
			c.JSON(200, gin.H{"user": claims})
		})
	}

	return r
}
