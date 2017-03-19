package main
import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "One"
	queue <- "Two"

	close(queue)

	for item := range queue {
		fmt.Println(item)
	}
}
