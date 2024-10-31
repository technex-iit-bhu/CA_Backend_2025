package utils

import (
	"CA_Backend/models"
	"fmt"
)

func GetReferralCode(user models.Users) string {
	ref := fmt.Sprintf("%s_ca_%s", user.Username, user.PhoneNumber[:5])
	return ref
}
