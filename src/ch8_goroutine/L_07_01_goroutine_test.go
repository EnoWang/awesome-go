package ch8_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// 0. goroutine -> Coroutine
/*
	Is a lighter thread
 */
func TestGoroutine(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+"goroutine %d\n", i)
			}
		}(i)
	}

	// note: make sure main function will not finish suddenly which will kill all goroutine
	time.Sleep(time.Millisecond)
}
