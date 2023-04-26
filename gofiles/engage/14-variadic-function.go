// Variadic functions can be called with any number of trailing arguments.

package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println("total: ", total)
}

type ss struct {
	l1 int
	l2 string
}

func su(su1 ...string) {
	fmt.Println(su1)
	for index, su1 := range su1 {
		if su1 == "cc" {
			fmt.Println("You are right the seq is ", index)
		}
	}
}
func main() {
	fmt.Println("here we see something different...")

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	su1 := make([]string, 10)
	su1 = append(su1, "cc", "caowenhao", "cc1", "cc2")
	su(su1...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
