package routes

import (
	"github.com/DeanDoyle1502/FYP-GigR.git/src/handlers"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	r *gin.Engine,
	userHandler *handlers.UserHandler,
	gigHandler *handlers.GigHandler,
	authHandler *handlers.AuthHandler,
	messageHandler *handlers.MessageHandler,
	chatSessionHandler *handlers.ChatSessionHandler,
	userRepo *repositories.UserRepository,
) {
	// Define user routes
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// Define gig routes
	gigs := r.Group("/gigs")
	gigs.Use(authHandler.Middleware())
	{
		gigs.POST("/", gigHandler.CreateGig)
		gigs.GET("/", gigHandler.GetAllGigs)
		gigs.GET("/mine", gigHandler.GetMyGigs)

		gigs.GET("/applications/mine", gigHandler.GetMyApplications)
		gigs.POST("/:gigID/apply", gigHandler.ApplyForGig)
		gigs.GET("/details/:id/applications", gigHandler.GetApplicationsForGig)
		gigs.POST("/:gigID/accept/:musicianID", gigHandler.AcceptMusicianForGig)

		gigs.POST("/:gigID/messages", messageHandler.SendMessage)
		gigs.GET("/:gigID/thread/:otherUserID", messageHandler.GetMessageThread)

		gigs.GET("/:gigID/session/:otherUserID", chatSessionHandler.GetOrCreateSession)
		gigs.PATCH("/:gigID/session/:otherUserID/complete", chatSessionHandler.MarkComplete)

		gigs.GET("/details/:id", gigHandler.GetGig)
		gigs.PUT("/details/:id", gigHandler.UpdateGig)
		gigs.DELETE("/details/:id", gigHandler.DeleteGig)
	}

	// Public gigs
	r.GET("/gigs/public", gigHandler.GetPublicGigs)

	// Ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Public auth routes
	r.POST("/auth/register", authHandler.RegisterUser)
	r.POST("/auth/login", authHandler.LoginUser)
	r.POST("/auth/confirm", authHandler.ConfirmUser)

	// Protected auth routes
	auth := r.Group("/auth")
	auth.Use(authHandler.Middleware())
	{
		auth.GET("/me", userHandler.GetCurrentUser)
	}
}
