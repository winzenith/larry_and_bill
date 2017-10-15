package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	routineAmount := 10
	wg.Add(routineAmount)

	// fun := func(i int) {
	// 	fmt.Println(i)
	// }
	// fun(100)

	for i := 0; i < routineAmount; i++ {

		// var iAddr *int
		// iAddr = &i

		// var iAddrAddr **int
		// iAddrAddr = &iAddr

		// **iAddrAddr === i
		// *iAddrAddr === &i
		// iAddrAddr === &&i
		// fmt.Println(**iAddrAddr) // == i

		go func(i int) {
			defer wg.Done()
			fmt.Println("sub goroutine", i)
		}(i)
	}

	wg.Wait() // block if wg is not 0
	fmt.Println("main goroutine")
}
