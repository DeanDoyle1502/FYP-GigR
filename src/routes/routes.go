package routes

import (
	"github.com/DeanDoyle1502/FYP-GigR.git/src/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, gigHandler *handlers.GigHandler) *gin.Engine {
	r := gin.Default()

	// Define user routes
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)

	// Define gig routes
	r.POST("/gigs", gigHandler.CreateGig)
	r.GET("/gigs", gigHandler.GetAllGigs)
	r.GET("/gigs/:id", gigHandler.GetGig)
	r.POST("/gigs/:id/apply", gigHandler.ApplyForGig)
	r.POST("/gigs/:gigID/accept/:musicianID", gigHandler.AcceptMusicianForGig)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
