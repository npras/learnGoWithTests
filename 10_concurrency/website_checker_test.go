package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	mockWebsiteChecker := func(url string) bool {
		if url == "http://whatthe.com" {
			return false
		}
		return true
	}
	websites := []string{
		"https://google.com",
		"https://bing.com",
		"http://whatthe.com",
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	want := map[string]bool{
		"https://google.com": true,
		"https://bing.com":   true,
		"http://whatthe.com": false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func BenchmarkCheckWebsitesWithChannel(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsitesWithChannel(slowStubWebsiteChecker, urls)
	}
}
