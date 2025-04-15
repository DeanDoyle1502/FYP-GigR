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
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// Define gig routes
	gigs := r.Group("/gigs")
	gigs.Use(middleware.AuthMiddleware())
	{
		gigs.POST("/", gigHandler.CreateGig)
		gigs.GET("/", gigHandler.GetAllGigs)
		gigs.GET("/:id", gigHandler.GetGig)
		gigs.GET("/mine", gigHandler.GetMyGigs)
		gigs.POST("/:id/apply", gigHandler.ApplyForGig)
		gigs.POST("/:id/accept/:musicianID", gigHandler.AcceptMusicianForGig)
		gigs.PUT("/:id", gigHandler.UpdateGig)
		gigs.DELETE("/:id", gigHandler.DeleteGig)

	}

	r.GET("/gigs/public", gigHandler.GetPublicGigs)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Authorisation
	r.POST("/auth/register", authHandler.RegisterUser)
	r.POST("/auth/login", authHandler.LoginUser)
	r.POST("/auth/confirm", authHandler.ConfirmUser)

	// Protected /auth routes (require valid JWT)
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/me", userHandler.GetCurrentUser)
	}

	return r
}
