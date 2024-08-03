package hello

import "fmt"

func Say(name string) {
	fmt.Println("Hello, ", name, "!")
}

func Accumulation(args ...int) int {
	var result int
	for _, v := range args {
		result += v
	}
	return result
}
