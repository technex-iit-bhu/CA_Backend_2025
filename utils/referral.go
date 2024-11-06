package utils

import (
	"CA_Backend/models"
	"fmt"
)

func GetReferralCode(user models.User) string {
	ref := fmt.Sprintf("%s_ca_%s", user.Username, user.CA_ID)
	return ref
}
