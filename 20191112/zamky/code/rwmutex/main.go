package mutex

// START OMIT
type RWMutex interface {
	// zápisové zámky
	Lock()
	Unlock()

	// čtecí zámky
	RLock()
	RUnlock()
}

// END OMIT
