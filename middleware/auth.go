package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/CircleConnectApp/post-service/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		
		cfg := config.LoadConfig()

		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			
			if userID, exists := claims["user_id"]; exists {
				c.Set("user_id", int(userID.(float64)))
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: missing user_id"})
				c.Abort()
				return
			}

			
			if role, exists := claims["role"]; exists {
				c.Set("role", role.(string))
			}

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
