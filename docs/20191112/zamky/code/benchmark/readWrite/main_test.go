package readWrite

import (
	"sync"
	"testing"
	"time"
)

type readWrite struct {
	sync.RWMutex
}

func (lock *readWrite) Read() {
	lock.RLock()
	defer lock.RUnlock()
	time.Sleep(time.Microsecond)
}

func BenchmarkReadWrite(b *testing.B) {
	lock := readWrite{}
	wg := sync.WaitGroup{}
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func() {
			defer wg.Done()
			lock.Read()
		}()
	}
	wg.Wait()
}
