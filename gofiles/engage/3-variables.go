/*
In Go, variables are explicitly declared and used by
the compiler to e.g. check type-correctness of function calls.
*/

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	fmt.Println("here we go again")

	first := sha256.New()
	//first.Write([]byte(ses.value))
	// var declares 1 or more variables.
	var a string = "This is Andela"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	// 0 for ints, 0.0 for floats, "" for strings, nil for pointers, …) We can also use the new function:
	// For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)
	type session struct {
		key   int
		value string
		name  int
		cache bool
	}
	var (
		ses session
		flo string
	)
	ses.value = "You are a hash man"
	flo = "Yes,you are right"
	// For example, the zero value for a string is ' '.
	var letter string
	fmt.Println(ses.value)
	fmt.Println(letter)
	first.Write([]byte(ses.value))
	fmt.Println(first.Sum(nil))
	second := sha512.New()
	second.Write([]byte(flo))
	first.Size()
	second.Size()
	fmt.Println(second.Sum(nil))
	// The := syntax is shorthand for declaring and initializing a variable, e.g.
	// for var f string = "short" in this case.
	short := "This is for short declaration."
	fmt.Println(short)
}
