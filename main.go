package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"
)

type statusCodes struct {
	m map[int]int
	o *sync.RWMutex
}

func newStatusCodes() *statusCodes {
	sr := new(statusCodes)

	sr.m = make(map[int]int)
	sr.m[http.StatusOK] = 0

	sr.o = new(sync.RWMutex)

	return sr
}

func (sr *statusCodes) increment(code int) {
	sr.o.Lock()
	defer sr.o.Unlock()

	sr.m[code] = sr.m[code] + 1
}

func (sr *statusCodes) String() string {
	var result string

	for code, count := range sr.m {
		result += fmt.Sprintf("Status code %d: %d\n", code, count)
	}

	return result
}

func main() {
	var url string
	var requests int
	var concurrency int
	var wg *sync.WaitGroup

	flag.StringVar(&url, "url", "https://www.fullcycle.com.br", "URL to send requests to")
	flag.IntVar(&requests, "requests", 100, "Number of requests to send")
	flag.IntVar(&concurrency, "concurrency", 10, "Number of requests to send concurrently")

	flag.Parse()

	fmt.Println("URL:", url)
	fmt.Println("Requests:", requests)
	fmt.Println("Concurrency:", concurrency)

	requested := 0
	wg = new(sync.WaitGroup)
	now := time.Now()
	codes := newStatusCodes()

	fmt.Println("Requesting...")
	for requested < requests {
		remainingRequests := requests - requested
		batch := int(math.Min(float64(concurrency), float64(remainingRequests)))
		wg.Add(batch)
		for j := 1; j <= batch; j++ {
			go func() {
				defer wg.Done()
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
				}
				codes.increment(resp.StatusCode)
				requested++
			}()
		}
		wg.Wait()
	}
	fmt.Println("Results:")
	fmt.Printf("All requests finished in %s\n", time.Since(now))
	fmt.Printf("Total requests: %d\n", requested)
	fmt.Println("Number of Responses By Status Code:")
	fmt.Println(codes)
}
