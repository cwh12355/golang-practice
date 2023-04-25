// Go supports constants of character, string, boolean, and numeric values.

package main

import (
	"fmt"
	"math"
)

// const declares a constant value.
const s string = "constant"

func main() {
	fmt.Println("here we discuss constants in golang!")
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const n = 500000
	const d = 3e20 / n
	// 定义const时可以显示或隐式类型
	const str string = "You are a bad man" //
	const Int int = 23
	const cc = "NO, I am a good woman"
	const bl = true
	const INT1 = 12
	const cons = 23.12223
	ss := Int + INT1

	fmt.Println("you can see there", ss)
	flo := math.Cos(cons)
	// Constant expressions perform arithmetic with arbitrary precision.
	fmt.Println(flo)

	// A numeric constant has no type until it’s given one, such as by an explicit cast.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))

}
