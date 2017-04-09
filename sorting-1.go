package main

import "fmt"
import "sort"

func main() {
	ints := []int{7, 3, 2, 8}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
