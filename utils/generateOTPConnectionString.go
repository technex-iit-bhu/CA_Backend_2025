package utils

func GenerateOTPConnectionString(username string) string {
	recoveryToken, _ := SerialiseRecovery(username)
	return recoveryToken
}
