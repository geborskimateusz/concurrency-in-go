package main

import (
	"fmt"
	"time"
)

// There might be 3 cases here:
// 1. Nothing is printed. Increment was executed before if statement
// 2. 0 is printed. If statement was executed before go routine
// 3. 1 is printed. If statement was executed then go routine increments value then value is printed
func FirstRaceCondition() {
	var data int

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("The value is %v", data)
	}
}

// Logically incorect. Althoutgh it may achieve some kind of correctnes.
// This is logically incorrect.
// time.Sleep can be used for debugging purpose but never used in real life scenarios.
func SecondRaceCondition() {
	var data int

	go func() {
		data++
	}()

	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("The value is %v", data)
	}
}
