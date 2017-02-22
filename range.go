package main

import "fmt" 

func main() {

	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum :", sum)
	
	dict := map[string]string{"a": "apple", "b": "banana"}

	for key, value := range dict {
		fmt.Println(key, value)
	}

	for key := range dict {
		fmt.Println(key)
	}

	for i, char := range "golang" {
		fmt.Println(i, char)
	}
}
