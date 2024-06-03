package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(count int) int {
		sum += count
		return sum

	}
}

func main() {
	counter := adder()
	fmt.Println(counter(10))
	fmt.Println(counter(10))
}
