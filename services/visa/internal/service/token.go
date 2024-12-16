package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	userIDClaim = "userID"
)

var secretKey = []byte("your-secure-secret-key")

type Token struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func createToken(userID int64, exp time.Duration) *jwt.Token {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			userIDClaim: userID,
			"exp":       time.Now().Add(exp).Unix(),
		})

	return token
}

func GenerateToken(userID int64) (string, string, error) {
	// Create the refresh token with user ID
	refreshToken := createToken(userID, time.Hour*24)

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	// Create the access token with user ID
	accessToken := createToken(userID, time.Minute*15)

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
