package main 
import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "string1"
	messages <- "string2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println("Out")
}

