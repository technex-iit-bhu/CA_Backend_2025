package utils

func GenerateRecoveryToken(username string) string {
	recoveryToken, _ := SerialiseRecovery(username)
	return recoveryToken
}
