package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"library-api/internal/config"
	"library-api/internal/handler"
	"library-api/internal/middleware"
	"library-api/internal/repository"
	"library-api/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	defer cfg.Close()

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable is not set")
	}

	// Initialize repository, service, and handler
	bookRepo := repository.New(cfg.Database)
	bookService := service.New(bookRepo)
	bookHandler := handler.NewHandler(bookService)

	userRepo := repository.NewUserRepository(cfg.Database)
	userService := service.NewUserService(userRepo, secretKey)
	userHandler := handler.NewUserHandler(userService, secretKey)

	router := gin.Default()

	// Public routes
	public := router.Group("/api")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
		}
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(secretKey))
	{
		books := protected.Group("/books")
		{
			books.POST("/", bookHandler.CreateBook)
			books.GET("/", bookHandler.GetBook)
			books.GET("/:id", bookHandler.GetAllBooks)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}
	}

	// Start the HTTP server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
