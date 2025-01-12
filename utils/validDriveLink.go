package utils

func IsValidDriveLink(link string) bool {

	if link == "" {
		return false
	}
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
