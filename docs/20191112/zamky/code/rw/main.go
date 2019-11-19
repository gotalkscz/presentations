package main

import (
	"sync"
)

// START OMIT
type RequestCounter struct {
	lock         sync.RWMutex // HL
	requestCount map[string]int
}

func (r *RequestCounter) Increment(method string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if count := r.requestCount[method]; count >= 500 {
		return
	}
	r.requestCount[method]++
}

func (r *RequestCounter) GetCount(method string) int {
	r.lock.RLock()         // HL
	defer r.lock.RUnlock() // HL
	return r.requestCount[method]
}

// END OMIT
