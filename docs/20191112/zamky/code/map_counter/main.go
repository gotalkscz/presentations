package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// START OMIT
type RequestCounter struct {
	lock         sync.Mutex
	requestCount map[string]int
}

func (r *RequestCounter) Increment(method string) { // HL
	r.lock.Lock()
	r.requestCount[method]++
	r.lock.Unlock()
}

func (r *RequestCounter) GetCount(method string) int {
	r.lock.Lock()
	count := r.requestCount[method]
	r.lock.Unlock()
	return count
}

// END OMIT

func main() {
	counter := RequestCounter{requestCount: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go func() {
			counter.Increment(http.MethodGet)
			// handle request
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			counter.Increment(http.MethodPost)
			// handle request
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println(http.MethodGet, counter.GetCount(http.MethodGet))
	fmt.Println(http.MethodPost, counter.GetCount(http.MethodPost))
}
