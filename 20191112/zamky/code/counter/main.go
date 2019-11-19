package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
type RequestCounter struct {
	lock         sync.Mutex // HL
	requestCount int
}

func (r *RequestCounter) Increment() {
	r.lock.Lock() // HL
	r.requestCount++
	r.lock.Unlock() // HL
}

func (r *RequestCounter) GetCount() int {
	r.lock.Lock() // HL
	count := r.requestCount
	r.lock.Unlock() // HL
	return count
}

// END OMIT

func main() {
	counter := RequestCounter{}
	for i := 0; i < 1000; i++ {
		go func() {
			counter.Increment()
			// handle request
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println(counter.GetCount())
}
