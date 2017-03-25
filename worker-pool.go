package main 

import "fmt"
import "time"

func worker(worker int, input <-chan int, output chan<- int) {

	for j := range input {
		fmt.Println("Worker ", worker, " started job ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker ", worker, " finished job ", j)
		output <- j * 2
	}
}

func main() {
	input := make(chan int, 100)
	output := make(chan int, 100)

	for i:=1; i<=100; i++ {
		go worker(i, input, output)
	}

	for j:=1; j<=100; j++ {
		input <- j
	}

	close(input)

	for i:=1; i<=100; i++{
		fmt.Println(<-output)
	}
}

