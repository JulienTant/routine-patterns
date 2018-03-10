package main

import (
	"fmt"
)

func main() {
	test := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			test <- i

		}
		close(test)
	}()

	for i := range test {
		fmt.Println(i, len(test), cap(test))
	}
}
