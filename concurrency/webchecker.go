package concurrency

import "fmt"

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

	fmt.Println("Begin iterating through urls...")
	r := <-resultsChannel
	/**
	Each time you pull an item from the channel using "[variable] := <- chan" syntax,
	it pops that item off the top of the resultsChannel
	*/
	fmt.Printf("From channel, '%s' r.url, '%v' r.valid\n", r.url, r.valid)
	results[r.url] = r.valid

	r = <-resultsChannel
	fmt.Printf("From channel, '%s' r.url, '%v' r.valid\n", r.url, r.valid)
	results[r.url] = r.valid

	r = <-resultsChannel
	fmt.Printf("From channel, '%s' r.url, '%v' r.valid\n", r.url, r.valid)
	results[r.url] = r.valid

	/**
	Can also iterate through the expected number of go routines to pull off their results
	This calls "r := <- resultsChannel" 3 times, which will pop the results.
	This only works because we know the exact number of results we're expecting.
	*/
	// for _, url := range url { // This isn't guaranteed to return the data in order it was called.
	// 	fmt.Printf("At %s\n", url)
	// 	r = <-resultsChannel
	// 	fmt.Printf("From channel, '%s' r.url, '%v' r.valid\n", r.url, r.valid)
	// 	results[r.url] = r.valid
	// }
	fmt.Println("Finished iterating through urls.")

	return results
}
