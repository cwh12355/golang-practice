// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it.

package main

import "fmt"

func main() {
	fmt.Println("here we go again bro! closures for real.")

	nextInt := intSeq()

	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each time we call nextInt.

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.

	newInts := intSeq()
	fmt.Println(newInts())
	seqfunc := seqsun()
	fmt.Println(seqfunc(256, "cc"))
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++

		return i
	}
}
func seqsun() func(int, string) (error, int, string) {
	i := 0
	str := []string{}
	return func(br1 int, inputstring string) (err error, in int, res string) {
		str = append(str, "cc1", "cc2", "cc3")
		i += 12
		br1 = 12
		for index, street := range str {
			if street == "cc" {
				br1 += 1
				fmt.Println(index)
				return nil, br1, street
				//fmt.Println(index)
			}
			if street == "ccr" {
				i += 1
				return nil, i, res
			}
		}
		return nil, br1, "You are a six star man"
	}
	//return nil, str[1]
}
