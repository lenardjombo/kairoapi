//github.com/lenardjombo/kairoapi/routes
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lenardjombo/kairoapi/internal/auth"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	auth := rg.Group("/auth")
	auth.POST("/register", handler.RegisterUser)
	auth.POST("/login", handler.LoginUser)
}
