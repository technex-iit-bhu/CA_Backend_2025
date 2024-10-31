package utils

func IsSafe(pwd string) bool {
	if len(pwd) < 8 {
		return false
	}

	// Checks if at least 1 special character, 1 upper case, and 1 number is present for safer password
	// ASCII values for special characters: 33-47, 58-64, 91-96, 123-126
	// ASCII values for upper case: 65-90
	// ASCII values for lower case: 97-122
	// ASCII values for numbers: 48-57
	return CheckAsciiLimit(pwd, 33, 126) && CheckAsciiLimit(pwd, 65, 90) && CheckAsciiLimit(pwd, 48, 57)
}

func CheckAsciiLimit(pwd string, ll int32, ul int32) bool {
	for _, c := range pwd {
		if c >= ll || c <= ul {
			return true
		}
	}
	return false
}
