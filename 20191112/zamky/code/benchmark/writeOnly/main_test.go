package writeOnly

import (
	"sync"
	"testing"
	"time"
)

type writeOnly struct {
	sync.Mutex
}

func (lock *writeOnly) Read() {
	lock.Lock()
	defer lock.Unlock()
	time.Sleep(time.Microsecond)
}

func BenchmarkWriteOnly(b *testing.B) {
	lock := writeOnly{}
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
