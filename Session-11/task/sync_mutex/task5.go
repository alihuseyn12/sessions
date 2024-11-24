package sync_mutex

import (
	"sync"
)

// ar counter int

func CheckMutex(countr *int, wg *sync.WaitGroup) {
	mu := sync.Mutex{}
	mu.Lock()
	*countr += 1
	mu.Unlock()
	wg.Done()
}
