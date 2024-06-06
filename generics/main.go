package main

import "fmt"

func getLast[T string | int](s []T) T {
	var defaultValue T
	if len(s) > 1 {
		return s[len(s)-1]
	}
	return defaultValue
}

func main() {
	value := getLast([]int{1, 2})

	fmt.Println(value)
}
