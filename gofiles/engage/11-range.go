// range iterates over elements in a variety of data structures.

package main

import "fmt"

func main() {
	fmt.Println("here we discuss golang range... Am I missing something?")

	// sum of numbers in a slice, arrays works like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	fmt.Println("sum: ", sum)

	for idx, val := range nums {
		if val == 3 {
			fmt.Println("index: ", idx)
		}
	}

	// range on map iterates over key/value pairs.

	// kvs := map[string]string{"a": "the fat boy jumped", "b": "where did he go"}
	// or
	kvs := make(map[string]string)
	kvs["a"] = "the fat boy jumped"
	kvs["b"] = "where did he go"

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range on strings iterates over Unicode code points.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	ac1 := []int{1, 2, 123}
	for index, exit := range ac1 {
		exit += 1
		fmt.Println("index is", index, exit)
		if exit == 2 {
			fmt.Println("exit value is :", index)
		}
	}
	fmt.Println()

}
