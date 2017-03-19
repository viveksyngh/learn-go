package main 
import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("message recieved : ", msg)
	default :
		fmt.Println("No message recieved")
	}

	msg := "hi"

	select {
	case messages <- msg :
		fmt.Println("Message sent : ", msg)
	default :
		fmt.Println("No meessage sent")
	}

	select {
	case msg := <- messages:
		fmt.Println("message recieved : ", msg)
	case signal := <-signals:
		fmt.Println("signals recieved : ", signal)
	default:
		fmt.Println("No Activity")
	}
}

