https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package main

import (
	"fmt"
	"time"
)

func increment(counter *int) {
	for i := 0; i < 1000000; i++ {
		*counter = *counter + 1
	}
}

func decrement(counter *int) {

	for i := 0; i < 1000000; i++ {
		*counter = *counter - 1
	}
}
func main() {

	counter := 0

	for i := 0; i < 1000; i++ {
		go increment(&counter)
	}

	for i := 0; i < 1000; i++ {
		go decrement(&counter)
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("Counter = %v", counter)
}
