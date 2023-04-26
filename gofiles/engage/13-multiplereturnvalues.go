//Go has built-in support for multiple return values. This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is where we go further")

	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	_, c := values()
	fmt.Println(c)

}

// The (int, int) in this function signature
// shows that the function returns 2 ints.

func values() (int, int) {
	return 7, 4
}

type first struct {
	bl1  bool
	flo  float32
	in1  int
	str  string
	por  *int
	por2 *float32
	por3 *string
}

var (
	fir first
	bl2 string
)

func light(light_piex int, length, wide int) (error, int, int) {
	area := length * wide
	volumes := *fir.por * wide
	fmt.Println("befor change area and volumes", area, volumes)
	area = *fir.por + *&fir.in1
	fmt.Println("after modify area", area)

}
