package main

import (
	"fmt"
	"time"
)

// - představme si REST API, kdy každý request spustí jednu gouroutine
// - chodí 1000 requestů za vteřinu

// START OMIT
type RequestCounter struct {
	requestCount int
}

func (r *RequestCounter) Increment() {
	r.requestCount++
}

func (r *RequestCounter) GetCount() int {
	return r.requestCount
}

func main() {
	counter := RequestCounter{}
	for i := 0; i < 1000; i++ { // <-- přijde 1000 requestů najednou
		go func() { // <-- odbavujeme paralelně
			counter.Increment()
			// handle request
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println(counter.GetCount())
}

// END OMIT
