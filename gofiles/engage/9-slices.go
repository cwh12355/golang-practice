package main

import "fmt"

// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

func main() {
	fmt.Println("Here we go again, something special this time")

	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	hex := make([]float32, 5)
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set ", s)
	fmt.Println("get ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	hex = append(hex, 12)
	fmt.Println("apd: ", s)
	fmt.Println("hex", hex)
	// Slices can also be copy’d.
	// Here we create an empty slice c of the same length as s and copy into c from s.

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].

	l := s[2:5]
	fmt.Println("sl1: ", l)

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2: ", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3: ", l)

	// We can declare and initialize a variable for slice in a single line as well.

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	c[0] = "facebook"
	c[1] = "google"
	c[3] = "yahoo"

	fmt.Println("set ", c)
	c = append(c, "twitter")

	fmt.Println(c)

}
