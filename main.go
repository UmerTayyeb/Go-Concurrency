package main

import (
	"fmt"
	"time"
)

func simpleFunction() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Simple function:", i)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

func goroutineFunction(j int, finish chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d: Goroutine: %d\n", j, i)
		time.Sleep(500 * time.Millisecond)
	}
	finish <- true // Signal that goroutine is finish
}

func main() {
	start := time.Now()
	finish := make(chan bool)
	for i := 0; i < 10; i++ {
		simpleFunction()
	}
	fmt.Println("simple func ended in: ", time.Since(start))

	start = time.Now()
	for i := 0; i < 10; i++ {
		go goroutineFunction(i, finish)
	}
	fmt.Println(time.Since(start))
	<-finish
	fmt.Println("go routine ended in: ", time.Since(start))
}
