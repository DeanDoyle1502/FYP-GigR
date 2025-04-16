package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwks *keyfunc.JWKS

func SetupJWKs() {
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	region := os.Getenv("AWS_REGION")

	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)

	var err error
	jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			fmt.Printf("Error refreshing JWKs: %v\n", err)
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to load Cognito JWKs: %v", err))
	}
}

// Middleware to verify JWT and attach Cognito sub to context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("🔐 AuthMiddleware triggered")

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			fmt.Println("❌ No Authorization header found")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			fmt.Println("❌ Invalid Authorization header format:", parts)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, jwks.Keyfunc)
		if err != nil || !token.Valid {
			fmt.Println("❌ Error parsing token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("❌ Error parsing claims from token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			fmt.Println("❌ Token is missing 'sub' field")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing user ID (sub)"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Set("sub", sub)

		c.Next()
	}
}
