package sync_waitgroup

import "sync"

func SumSlice(n []int, result *int, wg *sync.WaitGroup) {
	defer wg.Done()
	var count int
	for _, v := range n {
		count += v
		//fmt.Printf("value is %d", v)
	}
	*result = count
}
