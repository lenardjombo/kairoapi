package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lenardjombo/kairoapi/db/sqlc"
	"github.com/lenardjombo/kairoapi/internal/auth"
	"github.com/lenardjombo/kairoapi/pkg"
	"github.com/lenardjombo/kairoapi/routes"
)

func main() {
	// 1. Init DB from pkg/db.go
	if err := pkg.Init(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 2. Init SQLC Queries
	queries := db.New(pkg.DB)

	// 3. Init Repository → Service → Handler
	userRepo := auth.NewUserRepository(queries)
	authService := auth.NewAuthService(userRepo)
	authHandler := auth.NewHandler(authService)

	// 4. Settingup Gin & Auth Routes
	router := gin.Default()
	api := router.Group("/api")
	routes.RegisterAuthRoutes(api, authHandler)

	// 5. Starting Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server error:", err)
	}
}
