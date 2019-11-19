package main

import "sync"

// START OMIT
func main() {
	lock := sync.Mutex{}
	lock.Unlock() // <-- ajaj
}

// END OMIT
