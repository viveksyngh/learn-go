package main
import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	done := make(chan bool)

	go worker(done)

	<- done
}


