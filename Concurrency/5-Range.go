/*
Similar to slices and maps, channels can be ranged over.

for item := range ch {
    // item is the next value received from the channel
}

This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.
*/

/*
package main

import (
	"fmt"
	"time"
)

func concurrrentFib(n int) {
	ch := make(chan int)
	go fibonacci(n, ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// TEST SUITE - Don't touch below this line

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
*/