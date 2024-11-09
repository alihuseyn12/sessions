package main

func main() {
	//Task1
	//ch1 := make(chan string)
	//ch2 := make(chan string)
	//
	//go func() {
	//
	//	time.Sleep(time.Second * 1)
	//	ch1 <- "Hello from ch1"
	//}()
	//go func() {
	//
	//	time.Sleep(time.Second * 2)
	//	ch2 <- "Hello from ch2"
	//}()
	//select_statement.ExampleSelectStatment1(ch1, ch2)

	//task2

	//Task3
	//wg := sync.WaitGroup{}
	//
	//for i := 1; i <= 3; i++ {
	//	wg.Add(1)
	//	go sync_waitgroup.Task3(i, &wg)
	//}
	//wg.Wait()
	//fmt.Println("All goroutines have completed")

	//Task4
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//mid := len(slice) / 2
	//fristSlice := slice[:mid]
	//secndSlice := slice[mid:]
	//var sum1, sum2 int
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go sync_waitgroup.SumSlice(fristSlice, &sum1, &wg)
	//go sync_waitgroup.SumSlice(secndSlice, &sum2, &wg)
	//wg.Wait()
	//
	//fmt.Println("Partial sum 1:", sum1)
	//fmt.Println("Partial sum 2:", sum2)
	//fmt.Println("Total sum:", sum1+sum2)
	//Task5

	//wg := sync.WaitGroup{}
	//var counter int
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go sync_mutex.CheckMutex(&counter, &wg)
	//}
	//wg.Wait()
	//fmt.Println("Final counter value:", counter)

	//Task6
	//wg := sync.WaitGroup{}
	//wg.Add(3)
	//studentGr := map[string]int{}
	//go sync_mutex.AddStudentGr("Garay", 95, &studentGr, &wg)
	//go sync_mutex.AddStudentGr("Ali", 85, &studentGr, &wg)
	//go sync_mutex.AddStudentGr("Madina", 78, &studentGr, &wg)
	//wg.Wait()
	//fmt.Println("Final Grades:", studentGr)

	//task7

	//userData := make(map[string]int)

	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go sync_rwmutex.WrieDataUsers(&userData, &wg)
	//go sync_rwmutex.ReadData(userData, &wg)
	//wg.Wait()

	//task8

	//mut := sync.RWMutex{}
	//count := 0
	//sync_rwmutex.ReadCounter(count, &mut)
	//sync_rwmutex.IncrimmentCounter(&count, &mut)
	//sync_rwmutex.ReadCounter(count, &mut)
	//sync_rwmutex.IncrimmentCounter(&count, &mut)

	//task10 -----

}
