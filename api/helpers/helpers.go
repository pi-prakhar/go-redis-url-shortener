package helpers

import (
	"os"
	"regexp"
	"strings"
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url

}

func IsValidURL(url string) bool {
	// Regular expression pattern to match a URL
	pattern := `^((https?|ftp|file)://)?([a-zA-Z0-9.-]+)\.([a-zA-Z]{2,})(:[0-9]{2,5})?(/([a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=]*)?)?$`

	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)

	// Check if the input string matches the pattern
	return re.MatchString(url)
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
