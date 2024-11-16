package utils

import (
	"CA_Backend/models"
	"crypto/rand"
	"encoding/base32"
)

func GenerateCAID(user models.User) string {
	bytes := make([]byte, 10)
	rand.Read(bytes)
	id := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)[:16]
	return id
}
