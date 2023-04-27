package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	fmt.Println("here we discuss recursion in go")

	fmt.Println(fact(7))
	sell()

}

// This fact function calls itself until it reaches the base case of fact(0).

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func sell() int {
	//name := make([]string, 12)
	//cha := make(chan int, 3)
	ma := make(map[string]int, 1)
	ma["cass"] = 123
	//var area int
	for refIndex, mc := range ma {
		fmt.Println(refIndex, mc)
		if refIndex != "cass" {
			area := rand.Int() * rand.Int()
			vol := area * area
			fmt.Println(area, vol)
		}
	}
	return 0
}

// ma1 := make(map[string]int,20)
// ulint := make([]int,15)
func rand_string(length int) []string {
	const see = "ssssssdddfdsafdafadfadfd"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]string, 20)
	for i := range result {
		result[i] = string(see[seededRand.Int(len(see))]))
	}
	//ma1[name[2]] = ulint[5]

	return result

}
