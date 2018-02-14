package main

import ("fmt"
		"time"
		"sync")

var wg sync.WaitGroup

func do(s string) {
	defer wg.Done()
	for i := 0; i<3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	wg.Add(2)
	go do("one")
	go do("two")
	wg.Wait()
}
