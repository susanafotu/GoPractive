package main


import (
	"fmt"
	"sync"
)


type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}


type NoDupes struct {
	m   map[string]int
	mux sync.Mutex
}


var nodups = NoDupes{m: make(map[string]int)}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, pch chan int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.


	if depth <= 0 {
		pch <- 0
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		pch <- 0
		return
	}
	if val := nodups.m[url]; val <= 2 {
		fmt.Printf("found: %s %q\n", url, body)
	}


	cch := make(chan int)
	for _, u := range urls {
		nodups.mux.Lock()
		go Crawl(u, depth-1, fetcher, cch)
		nodups.m[u]++
		nodups.mux.Unlock()
	}


	for i := 0; i < len(urls); i++ {
		<-cch
	}
	pch <- 0
	return
}


func main() {
	ch := make(chan int)
	go Crawl("http://golang.org/", 4, fetcher, ch)
	<-ch
}


// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult


type fakeResult struct {
	body string
	urls []string
}


func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}


// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

