package main

import (
	"golang-clean-crud/config"
	"golang-clean-crud/handler"
	"golang-clean-crud/models"
	"golang-clean-crud/repository"
	"golang-clean-crud/routes"
	"golang-clean-crud/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Product{})

	r := gin.Default()

	// Dependency Injection
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo)
	productsHandler := handler.NewProductHandler(productService)

	routes.SetupRoutes(r, productsHandler)

	r.Run(":8080")
}
