package utils

import (
	"CA_Backend/config"
	"fmt"
	"time"
	"log"
	"github.com/golang-jwt/jwt/v5"
)

var serialKey = []byte(config.Config("JWT_SECRET"))
var recoveryKey = []byte(config.Config("RECOVERY_SECRET"))

func SerialiseUser(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	signedToken, err := token.SignedString(serialKey)
	if err != nil {
		return "", err
	}
	log.Printf("Deserialising token: %+v\n", signedToken)

	return signedToken, nil
}

func DeserialiseUser(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return serialKey, nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token or claims")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().After(time.Unix(int64(exp), 0)) {
			return "", fmt.Errorf("token expired")
		}
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username claim is missing or not a string")
	}
	log.Printf("Token claims: %+v\n", claims)

	return username, nil
}

func SerialiseRecovery(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"expires_at": time.Now().Add(10 * time.Minute).Unix(),
	})
	signedToken, err := token.SignedString(recoveryKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseRecovery(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(recoveryKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	if time.Now().After(time.Unix(int64(claims["expires_at"].(float64)), 0)) {
		return "", fmt.Errorf("expired")
	}
	return claims["username"].(string), nil
}
