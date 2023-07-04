package concurrency

type WebsiteChecker func(url string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, websites []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, website := range websites {
		go func(w string) {
			// Send statement
			resultChannel <- result{w, wc(w)}
		}(website)
	}

	for i := 0; i < len(websites); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
