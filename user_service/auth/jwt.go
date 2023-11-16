package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"time"
)

func GenerateToken(userId uint, secretKey []byte, duration int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(time.Duration(duration) * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Error generating signed string.")
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey []byte) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Error while parsing token.")
		return 0, err
	}

	if !token.Valid {
		log.Info().Msg("Token is not valid!")
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Info().Msg("Couldn't parse claims!")
		return 0, errors.New("couldn't parse claims")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		log.Info().Msg("userId claim is not of expected type!")
		return 0, errors.New("userId claim is not of expected type")
	}

	return uint(userIdFloat), nil
}
