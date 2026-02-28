package handler

import (
	"golang-clean-crud/dto"
	"golang-clean-crud/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Register(req)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Register success"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	c.ShouldBindJSON(&req)

	token, err := h.service.Login(req)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login success",
		"token":   token,
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	c.ShouldBindJSON(&body)

	token, err := h.service.RefreshToken(body.RefreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid Refresh Token"})
		return
	}

	c.JSON(200, gin.H{"access_token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// Get user_id from JWT Middleware
	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID := uint(userIDValue.(float64))

	err := h.service.Logout(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Logout success",
	})
}
