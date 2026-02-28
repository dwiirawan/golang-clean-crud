package routes

import (
	"golang-clean-crud/handler"
	"golang-clean-crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	productHandler *handler.ProductHandler,
	authHandler *handler.AuthHandler,
) {
	// PUBLIC ROUTES
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}

	// PROTECTED ROUTES (JWT)
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// PRODUCT CRU
		api.GET("/products", productHandler.GetAll)
		api.GET("/products/:id", productHandler.GetByID)
		api.POST("/products", productHandler.Create)
		api.PUT("/products/:id", productHandler.Update)

		// LOGOUT
		api.POST("/logout", authHandler.Logout)
	}

	// PROTECTED ADMIN ONLY
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.DELETE("/products/:id", productHandler.Delete)
	}
}
