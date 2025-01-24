package concurrency

type WebsiteChecker func(string) bool
type result struct {
	url   string
	valid bool
}

func CheckWebsites(wc WebsiteChecker, url []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range url {
		go func() {
			resultsChannel <- result{url: url, valid: wc(url)}
		}()
	}

	for i := 0; i < len(url); i++ {
		r := <-resultsChannel
		results[r.url] = r.valid
	}

	return results
}
