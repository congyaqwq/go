package main

import "fmt"

// 1
type person struct {
	name string
	age int
}

// 2
func newPerson(name string, age int) *person {
	p := person{name: name, age: age}

	return &p
}

// 嵌套
type man struct {
	person
	work string
}

func main() {
	fmt.Println(person{name: "foo", age: 18})

	s := person{name: "bar", age: 19}

	sp := &s

	fmt.Println(sp, sp.name)

	// 3
	dog := struct {
		name string
	}{
		name: "Teddy",
	}
	fmt.Println(dog, &dog)


	// 嵌套
	foo := man{
		person: person{
			age: 1,
			name: "Jack",
		},
	}

	fmt.Println(foo, "foo")
}