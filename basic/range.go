package main

import "fmt"

func main() {
	var nums = []int{2,3,4}
	sum := 0
	for _,num := range nums {
		sum += num
	}

	for index := range nums {
		if(index==2) {
			fmt.Println(("2"))
		}
	}

	for i, c := range "go" {
		// string -> The first value is the starting byte index of the rune and the second the rune itself.
		fmt.Println(i, c) // 1 111
	}
}