package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// chceme rozdělit jednotlivé metody requestů, proto si na místo int dáme mapu string int
// skryl jsem tu jednu záludnost

// START OMIT
type RequestCounter struct {
	lock         sync.Mutex
	requestCount map[string]int // HL
}

func (r RequestCounter) Increment(method string) { // HL
	r.lock.Lock()
	r.requestCount[method]++ // HL
	r.lock.Unlock()
}

func (r *RequestCounter) GetCount(method string) int { // HL
	r.lock.Lock()
	count := r.requestCount[method] // HL
	r.lock.Unlock()
	return count
}

// END OMIT

// START B OMIT
func main() {
	counter := RequestCounter{requestCount: make(map[string]int)} // <-- inicializace mapy // HL
	for i := 0; i < 1000; i++ {
		go func() {
			counter.Increment(http.MethodGet) // HL
			// handle GET request
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			counter.Increment(http.MethodPost) // HL
			// handle POST request
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println(http.MethodGet, counter.GetCount(http.MethodGet))   // HL
	fmt.Println(http.MethodPost, counter.GetCount(http.MethodPost)) // HL
}

// END B OMIT
