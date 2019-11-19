package main

import "sync"

// START OMIT
func main() {
	lock := sync.Mutex{}
	lock.Lock()
	lock.Lock() // <-- ajaj
}

// END OMIT
