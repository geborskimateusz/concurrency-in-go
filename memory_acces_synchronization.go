package main

import (
	"fmt"
	"sync"
)

// Assume that code between Lock and Unclock have exclusive acces
func SynchronizedMemoryAccessFirst() {
	var memoryAccess sync.Mutex
	var value int

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("The value is 0 \n")
	} else {
		fmt.Printf("The value is %v", value)
	}
	memoryAccess.Unlock()
}

func SynchronizedMemoryAccessSec() {
	var value int

	ints := make(chan int)

	go func(ch chan int) {
		defer close(ch)
		value++
		ints <- value
	}(ints)

	for integer := range ints {
		if value == 0 {
			fmt.Printf("The value is 0 \n")
		}
		fmt.Println("received", integer)
	}
}
