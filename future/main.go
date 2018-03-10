package main

import (
	"fmt"
	"time"
)

func asyncData() <-chan string {
	// we can set the size of the channel to 1
	c := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		c <- "from the future"
		close(c)
	}()

	return c

}

func main() {
	c := asyncData()

	fmt.Println("we can still do some stuff instead of just waiting")

	// Then we get the data, or wait for it to be ready
	data := <-c

	fmt.Printf("Got '%s'", data)
}
