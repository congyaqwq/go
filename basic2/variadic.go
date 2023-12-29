package main

import "fmt"

func sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1,2)
	a := []int{1,2}
	sum(a...)
}