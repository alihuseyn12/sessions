package sync_rwmutex

import (
	"fmt"
	"sync"
)

func IncrimmentCounter(counter *int, mut *sync.RWMutex) {
	mut.Lock()
	*counter++
	fmt.Println("Writer updated counter:", *counter)
	defer mut.Unlock()
}
func ReadCounter(counter int, mut *sync.RWMutex) {
	mut.RLock()
	fmt.Println("Reader accessed counter:", counter)
	defer mut.RUnlock()

}
