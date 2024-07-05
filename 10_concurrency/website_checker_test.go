package concurrency

import (
	"testing"
	//"reflect"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://whatthe.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites = []string{
		"https://google.com",
		"https://bing.com",
		"http://whatthe.com",
	}
}
