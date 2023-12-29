package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["1"] = 1
	m["2"] = 2

	fmt.Println(m, len(m)) // map[1:1 2:2] 2

	delete(m, "1")
	fmt.Println(m, len(m)) // map[2:2] 1

	clear(m)
	fmt.Println(m, len(m)) // map[] 0

	t, prs := m["2"]
	// t -> 值，没有默认返回 zero value, prs -> bool 是否存在
	fmt.Println(t, prs) // 0 false

	n := map[string]int{"1": 1, "2": 2}
	n2 := map[string]int{"1": 1, "2": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2") // n == n2
	}
}
