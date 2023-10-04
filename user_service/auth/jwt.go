package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user_service/models"
)

func GenerateToken(user *models.User, secretKey []byte, duration int) (string, error) {
	username := user.Username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Duration(duration) * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey []byte) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	idValue, ok := claims["id"]
	if !ok {
		return nil, errors.New("id claim not found")
	}

	id, ok := idValue.(float64) // JWT numeric values are decoded as float64
	if !ok {
		return nil, errors.New("id claim is not a uint")
	}

	user := &models.User{
		ID:       uint(id),
		Username: token.Claims.(jwt.MapClaims)["username"].(string),
	}

	print(user)
	print(token.Claims.(jwt.MapClaims)["id"].(string))
	print(token.Claims.(jwt.MapClaims)["username"].(string))

	return user, nil
}
