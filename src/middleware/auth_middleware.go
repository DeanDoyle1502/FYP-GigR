package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwks *keyfunc.JWKS

func SetupJWKs() {
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	region := os.Getenv("AWS_REGION")

	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
	fmt.Println("🔐 Loading Cognito JWKs from:", jwksURL)

	var err error
	jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			fmt.Printf("❌ Error refreshing JWKs: %v\n", err)
		},
	})
	if err != nil {
		panic(fmt.Sprintf("❌ Failed to load Cognito JWKs: %v", err))
	}
	fmt.Println("✅ JWKs loaded successfully")
}

// Middleware to verify JWT and attach Cognito sub to context
func AuthMiddleware(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("🔐 AuthMiddleware triggered")

		authHeader := c.GetHeader("Authorization")
		fmt.Println("🛡️  Received Authorization header:", authHeader)

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
		fmt.Println("🔐 Token string extracted")

		token, err := jwt.Parse(tokenString, jwks.Keyfunc)
		if err != nil {
			fmt.Println("❌ jwt.Parse() failed:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("❌ Token is invalid")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		fmt.Println("✅ Token is valid")

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("❌ Failed to cast claims from token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			fmt.Println("❌ Token is missing 'sub' field:", claims)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing user ID (sub)"})
			c.Abort()
			return
		}

		fmt.Println("✅ Token claims extracted. sub =", sub)

		c.Set("user", claims)
		c.Set("userID", sub) // make sure handler reads 'userID' key

		c.Next()
	}
}
