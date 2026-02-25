package routes

import (
	"golang-clean-crud/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.ProductHandler) {
	api := r.Group("/api")

	api.GET("/products", h.GetAll)
	api.GET("/products/:id", h.GetByID)
	api.POST("/products", h.Create)
	api.PUT("/products/:id", h.Update)
	api.DELETE("/products/:id", h.Delete)
}
