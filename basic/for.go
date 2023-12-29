package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for n := 0; n <= 2; n++ {
		fmt.Println(n)
	}
}