package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET")) // atau langsung tulis []byte("secret123")

// GenerateToken membuat JWT token dari user ID
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ✅ Validasi token dan kembalikan user ID
func ValidateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Validasi metode signing
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// ✅ Ambil user ID dari claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	return int(userIDFloat), nil
}
