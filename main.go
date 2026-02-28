package main

import (
	"golang-clean-crud/config"
	"golang-clean-crud/entity"
	"golang-clean-crud/handler"
	"golang-clean-crud/models"
	"golang-clean-crud/repository"
	"golang-clean-crud/routes"
	"golang-clean-crud/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&entity.User{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Dependency Injection
	// repository
	productRepo := repository.NewProductRepository(config.DB)
	userRepo := repository.NewUserRepository(config.DB)

	// service
	productService := service.NewProductService(productRepo)
	userService := service.NewAuthService(userRepo)

	// handler
	productsHandler := handler.NewProductHandler(productService)
	authHandler := handler.NewAuthHandler(userService)

	routes.SetupRoutes(r, productsHandler, authHandler)

	r.Run(":8080")
}
