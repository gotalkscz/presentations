package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// START OMIT
type RequestCounter struct {
	requestCount int32
}

func (r *RequestCounter) Increment() {
	atomic.AddInt32(&r.requestCount, 1) // HL
}
func (r *RequestCounter) GetCount() int32 {
	return atomic.LoadInt32(&r.requestCount) // HL
}
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

// END OMIT
