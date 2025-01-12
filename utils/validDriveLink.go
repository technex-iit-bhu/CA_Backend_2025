package utils

func IsValidDriveLink(link string) bool {
	// Check if link is empty
	if link == "" {
		return false
	}

	// Check if link starts with Google Drive URL patterns
	validPrefixes := []string{
		"https://drive.google.com/",
		"https://docs.google.com/",
		"http://drive.google.com/",
		"http://docs.google.com/",
	}

	isValid := false
	for _, prefix := range validPrefixes {
		if len(link) >= len(prefix) && link[:len(prefix)] == prefix {
			isValid = true
			break
		}
	}

	return isValid
}