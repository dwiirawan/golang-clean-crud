package handler

import (
	"golang-clean-crud/models"
	"golang-clean-crud/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{s}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	data, _ := h.service.GetAll()
	c.JSON(http.StatusOK, data)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.service.GetbyID(uint(id))

	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, data)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var input models.Product
	c.ShouldBindJSON(&input)

	data, err := h.service.Create(input)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Product
	c.ShouldBindJSON(&input)

	data, err := h.service.Update(id, input)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.service.Delete(uint(id))

	c.JSON(200, gin.H{"message": "deleted"})
}
