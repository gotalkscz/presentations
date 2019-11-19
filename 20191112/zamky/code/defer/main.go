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

func (r *RequestCounter) Increment(method string) {
	r.lock.Lock()
	defer r.lock.Unlock() // HL
	if count := r.requestCount[method]; count >= 500 {
		return
	}
	r.requestCount[method]++
}

func (r *RequestCounter) GetCount(method string) int {
	r.lock.Lock()
	defer r.lock.Unlock() // HL
	return r.requestCount[method]
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
