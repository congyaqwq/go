package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("福来 day")
	default:
		fmt.Println("work")
	}

	t := time.Now().Weekday()
	switch {
	case t == time.Friday:
		fmt.Println("福来 day")
	default:
		fmt.Println("work")
	}

	whichType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("no type")
		}
	}

	whichType(true)
}
