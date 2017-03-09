package main

import "fmt"

type rect struct {
	width int
	height int 
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * ( r.width + r.height)
}

func main() {
	r := rect{10, 8}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r

	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}


