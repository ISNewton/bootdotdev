package main

import "fmt"

func main() {
	fmt.Println(bulkSend(4))
}
func bulkSend(numMessages int) float64 {
	total := 0.0
	for i := 0; i <= numMessages; i++ {
		total += 1 + float64(i)/100
		fmt.Println(1, float64(i)/100)
	}
	return total
}
