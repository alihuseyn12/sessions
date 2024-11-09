package sync_rwmutex

import (
	"fmt"
	"sync"
	"time"
)

func WrieDataUsers(user *map[string]int, wg *sync.WaitGroup) {
	mux := sync.RWMutex{}
	mux.RLock()
	(*user)["Garay"] = 25
	(*user)["Ali"] = 25
	(*user)["Madina"] = 28
	mux.RUnlock()
	defer wg.Done()
}
func ReadData(user map[string]int, wg *sync.WaitGroup) {
	mux := sync.RWMutex{}
	mux.RLock()
	fmt.Println(user)
	mux.RUnlock()
	defer wg.Done()
	time.Sleep(time.Second * 1)

}
