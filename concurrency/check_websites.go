package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

// slow version of website checker benchmark
// goos: darwin
// goarch: arm64
// pkg: github.com/eclecticjohny/learn-go-with-tests/concurrency
// BenchmarkCheckWebsites-8   	       1	2086890375 ns/op	     112 B/op	       4 allocs/op
// PASS
// ok  	github.com/eclecticjohny/learn-go-with-tests/concurrency	2.415s
