package main

import "fmt"

type rect struct {
	width, height int
}

func area(r *rect) int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 12}

	fmt.Println(area(&r))

	rp := &r

	fmt.Println(rp.perim())
	fmt.Println(r.perim())
}
