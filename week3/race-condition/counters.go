https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func increment(counter *int, atomicValue *int32, group *sync.WaitGroup) {

	for !atomic.CompareAndSwapInt32(atomicValue, 0, 1) {
		runtime.Gosched() //allow other goroutines to do stuff.
	}
	// ---------------------------- Start of the Critical Section
	for i :=0; i < 1000000; i++ {
		*counter = *counter + 1
	}
	// ---------------------------- End of the Critical Section
	atomic.StoreInt32(atomicValue, 0)
	group.Done()
}

func decrement(counter *int, atomicValue *int32, group *sync.WaitGroup) {


	for !atomic.CompareAndSwapInt32(atomicValue, 0, 1) {
		runtime.Gosched() //allow other goroutines to do stuff.
	}
	// ---------------------------- Start of the Critical Section
	for i :=0; i < 1000000; i++ {
		*counter = *counter - 1
	}
	// ---------------------------- End of the Critical Section
	atomic.StoreInt32(atomicValue, 0)
	group.Done()
}
func main() {

	counter := 0
	var atomicValue int32
	var group sync.WaitGroup


	group.Add(1000)
	for i := 0; i < 1000; i++ {
		go increment(&counter,&atomicValue, &group)
	}

	group.Add(1000)
	for i := 0; i < 1000; i++ {
		go decrement(&counter,&atomicValue, &group)
	}

	group.Wait()
	fmt.Printf("Counter = %v", counter)
}
