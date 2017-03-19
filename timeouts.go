package main 
import "fmt"
import "time"

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result"
	}()

	select {
	case msg := <- c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "result"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 2")
	}
}

