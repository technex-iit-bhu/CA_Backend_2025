package utils

import (
	"fmt"
	"technexRegistration/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var serialKey = []byte(config.Config("JWT_SECRET"))
var gmailKey = []byte(config.Config("GMAIL_SECRET"))
var githubKey = []byte(config.Config("GITHUB_SECRET"))
var recoveryKey = []byte(config.Config("RECOVERY_SECRET"))

func SerialiseUser(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	signedToken, err := token.SignedString(serialKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseUser(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(serialKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["username"].(string), nil
}

func SerialiseGmailToken(gmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"gmail": gmail,
	})
	signedToken, err := token.SignedString(gmailKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseGmailToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(gmailKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["gmail"].(string), nil
}

func SerialiseGithubToken(github string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"github": github,
	})
	signedToken, err := token.SignedString(githubKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseGithubToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(githubKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["github"].(string), nil
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
