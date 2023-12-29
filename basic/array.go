package main

import "fmt"

func main() {
	var a  = [5]int{1,1,1,1,1}
	var b [5] int 
	fmt.Println(a, b)
	// length
	fmt.Println(len(a))
	// 二维数组
	var c [2][3]int

	fmt.Println(c)
}