package main

import "fmt"

// 参数类型不写会默认为最后一个的类型
func add(a, b, c int) int {
	return a + b + c
}

func multiple(a, b int) (int, int) {
	return a, b
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	a := add(1, 2, 3)
	b, c := multiple(1, 2)
	d := sum(1,2,3,4)
	fmt.Println(a, b, c, d)
}
