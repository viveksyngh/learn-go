package main
import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer1 has expired")

	timer2 := time.NewTimer(time.Second * 2)

	go func() {
		<-timer2.C
		fmt.Println("Timer2 has expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timers 2 has stopped")
	}
}
