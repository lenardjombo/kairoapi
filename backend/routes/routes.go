//github.com/lenardjombo/kairoapi/routes
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenardjombo/kairoapi/internal/auth"
	"github.com/lenardjombo/kairoapi/internal/middleware" 
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	// Public routes
	authRoutes := rg.Group("/auth")
	authRoutes.POST("/register", handler.RegisterUser)
	authRoutes.POST("/login", handler.LoginUser)

	// Protected routes group
	protectedRoutes := rg.Group("/protected")
	protectedRoutes.Use(middleware.AuthMiddleware()) 
	{
		// Example protected route
		// api/protected/profile - demo for testing protected route
		protectedRoutes.GET("/profile", func(c *gin.Context) {
			userID, exists := c.Get("userID")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
				return
			}
			username, exists := c.Get("username")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Username not found in context"})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message":  "Welcome to the protected demo route",
				"userID":   userID,
				"username": username,
			})
		})
	}
}