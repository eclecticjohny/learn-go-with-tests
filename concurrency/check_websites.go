package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
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

// non-concurrent version of website checker benchmark
// goos: darwin
// goarch: arm64
// pkg: github.com/eclecticjohny/learn-go-with-tests/concurrency
// BenchmarkCheckWebsites-8   	       1	2086890375 ns/op	     112 B/op	       4 allocs/op
// PASS
// ok  	github.com/eclecticjohny/learn-go-with-tests/concurrency	2.415s

// concurrent version of website checker benchmark
// goos: darwin
// goarch: arm64
// pkg: github.com/eclecticjohny/learn-go-with-tests/concurrency
// BenchmarkCheckWebsites-8   	      56	  21230214 ns/op	   17697 B/op	     319 allocs/op
// PASS
// ok  	github.com/eclecticjohny/learn-go-with-tests/concurrency	1.307s
