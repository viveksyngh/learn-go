package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {

	return a + b + c
}

func main() {
	res := plus(2, 3)
	fmt.Println("res : ", res)

	res = plusPlus(2, 3, 4)
	fmt.Println("res : ", res)
}

