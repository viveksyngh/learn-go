package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs 
			if more {
				fmt.Println("Job recieved : ", job)
			} else {
				fmt.Println("All Jobs receievd.")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sent job : ", i)
	}
	close(jobs)
	fmt.Println("Sent all jobs")
	<-done
}



