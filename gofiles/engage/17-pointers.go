// Go supports pointers, allowing you to pass references to values and records within your program.

package main

import "fmt"

func main() {
	fmt.Println("Go pointers explained here.")

	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	// Pointers can be printed too.
	fmt.Println("pointer: ", &i)

	// zeroval doesn’t change the i in main,
	// but zeroptr does because it has a reference to the memory address for that variable.
	j := "xiaojiji"
	ac1ptr(&j)
	fmt.Println("gne_ptr value and locatin :", j, &j)

}

// zeroval has an int parameter, so arguments will be passed to it by value not by address.
// zeroval will get a copy of ival distinct from the one in the calling function.

func zeroval(ival int) {
	ival = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory address
// to the current value at that address.
// Assigning a value to a dereferenced pointer changes the value at the referenced address.

// further note:: When you want to access the data/value in the memory that the pointer points to -
// the contents of the address with that numerical index - then you dereference the pointer.

// here iprt is of type integer pointer
func zeroptr(iprt *int) {
	*iprt = 0
}
func ac1ptr(gen_str *string) {
	*gen_str = "xiaojingjing dead"
}
