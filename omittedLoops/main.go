package main

import "fmt"

func main() {
	fmt.Println(maxMessages(220))
}

func maxMessages(thresh int) int {
	totalCost := 0
	for i := 0; ; i++ {
		totalCost += 100 + i
		if totalCost >= thresh || thresh < 100 {
			return i
		}
	}
}
