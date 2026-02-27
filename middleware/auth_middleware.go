package middleware

import (
	"golang-clean-crud/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := jwt.Parse(tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return utils.ACCESS_SECRET, nil
			})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()
	}
}
