/*
Go’s structs are typed collections of fields.
They’re useful for grouping data together to form records.
*/

package main

import (
	"fmt"
)

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}
type c struct {
	name     string
	age      int
	color    string
	good_man bool
}

func main() {
	fmt.Println("let's go a little bit further...")

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 23})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Andela"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	var cc1 = c{"caocao", 22, "yellow", true}
	fmt.Println(cc1)

}

/*type cc struct {
	name string
	age int
	color string
	good_man bool
 } */
