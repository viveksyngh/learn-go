package main

import "fmt"
import "math"

func main() {
	const n = 500000000

	const b = 3e20 / n
	fmt.Println(b)
	fmt.Println(int64(b))
	fmt.Println(math.Sin(n))
}

