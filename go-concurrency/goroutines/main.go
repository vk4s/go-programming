package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"epsilon",
		"zeta",
	}

	// add the number of items to wait for
	// use legth of slice/array for a dynamic number of items
	wg.Add(len(words))

	for i, v := range words {
		// pass a "REFERENCE" of the waitgroup object
		go printSomething(fmt.Sprintf("%d: %s", i, v), &wg)
	}
	// time.Sleep(1 * time.Second)

	wg.Add(1)
	go printSomething("Second go routine", &wg)

	// Wait till the go routine execution finishes
	wg.Wait()
}

func printSomething(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	// Mark the execution as done.
	// if you mark it before running all your code then
	// there is no gaurantee that the code exdecuted after this point will be accessible
	// and your main routine will wait for the response. It will (main routine) finish early
	// and your program will exit.
	// SO, ALWAYS PUT wg.Done() after all the statements are done executing in your go routine.
	wg.Done()
}
