package main

import "fmt"

type result struct {
	workerID int
	original int
	result   int
}

func sqrtWorker(ID int, in <-chan int, out chan<- result) {
	for i := range in {
		out <- result{workerID: ID, original: i, result: i * i}
	}
}

func main() {

	nbWorker := 5
	nbSqrt := 100

	in := make(chan int)
	out := make(chan result)

	// Start the workers
	for i := 0; i < nbWorker; i++ {
		go sqrtWorker(i+1, in, out)
	}

	// generate the data to process in the background
	go func() {
		for i := 0; i <= nbSqrt; i++ {
			in <- i
		}
		close(in)
	}()

	// get the result
	for i := 0; i <= nbSqrt; i++ {
		res := <-out
		fmt.Printf("[worker %d] sqrt of %d is %d\n", res.workerID, res.original, res.result)
	}

}
