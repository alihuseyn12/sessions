package sync_mutex

import (
	"sync"
)

func AddStudentGr(name string, grade int, m *map[string]int, wg *sync.WaitGroup) {
	mu := sync.Mutex{}
	mu.Lock()
	(*m)[name] = grade
	wg.Done()
	mu.Unlock()

}
