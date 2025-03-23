package routes

import (
	"fmt"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, gigHandler *handlers.GigHandler, authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "home"})
	})

	fmt.Println("🔥 ROUTES: SetupRouter() loaded")

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

	r.GET("/debug/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Auth route group is active"})
	})

	return r
}
