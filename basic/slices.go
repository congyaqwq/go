package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	// nil -> init zero value
	fmt.Println(s == nil) // true

	s = make([]string, 3)
	// cap -> for Slices: the maximum length the slice can reach when resliced;
	fmt.Println(s, len(s), cap(s))

	s = append(s, "d", "e")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// 类似 slice 包头不包尾
	l := c[1:2]
	fmt.Println(l)

	t := []string{"a", "b", "c"}

	if slices.Equal(l, t) {
		// do something
	}

	// 二维数组
	twoS := make([][]int, 2)
	fmt.Println(twoS)
}
