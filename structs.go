package main

import "fmt"

type person struct {
	name string 
	age int
}

func main() {
	fmt.Println(person{"Vivek", 26})

	fmt.Println(person{name: "Vivek", age: 26})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})
	
	s := person{name: "Vivek", age: 40}
	fmt.Println(s.name, s.age)
	
	sp := &s
	sp.age = 60

	fmt.Println(sp.age)
	fmt.Println(s.age)
}
