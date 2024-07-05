package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

func CheckWebsitesWithChannel(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
