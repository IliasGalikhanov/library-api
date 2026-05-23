package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"library-api/internal/config"
	"library-api/internal/handler"
	"library-api/internal/repository"
	"library-api/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	defer cfg.Close()

	// Initialize repository, service, and handler
	repo := repository.New(cfg.Database)
	srv := service.New(repo)
	h := handler.NewHandler(srv)
	router := gin.Default()

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("", h.CreateBook)
			books.GET("/:id", h.GetBook)
			books.GET("", h.GetAllBooks)
			books.PUT("/:id", h.UpdateBook)
			books.DELETE("/:id", h.DeleteBook)
		}
	}

	// Start the HTTP server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
