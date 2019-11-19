package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	rwLock := sync.RWMutex{}
	number := 5
	rwLock.RLock()   // čtecí zámek
	number = 10      // GO tomuto nezabrání // HL
	rwLock.RUnlock() // odemknutí čtecího zámku
	fmt.Println(number)
}

// END OMIT
