package utils

import (
	"CA_Backend/models"
	"github.com/teris-io/shortid"
)

func GenerateCAID(user models.User) string {
	uniqueID, _ := shortid.Generate()
	return uniqueID
}
