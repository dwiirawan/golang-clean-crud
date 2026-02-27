package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ACCESS_SECRET = []byte("ACC_Secret123")
var REFRESH_SECRET = []byte("REF_Secret123")

func GenerateAccessToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(ACCESS_SECRET)
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(REFRESH_SECRET)
}

func VerifyRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// pastikan algoritma benar
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return REFRESH_SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	// validasi token
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	// ambil claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid claims")
	}

	return claims, nil
}
