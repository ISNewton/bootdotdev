package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	// ?
	costs := make([]float64, len(messages))

	for i, message := range messages {
		costs[i] = float64(len(message)) * 0.01
	}

	return costs
}

func main() {

	var messages []string
	for i := 0; i < 1000000; i++ {
		messages = append(messages, "hello")
	}
	fmt.Println(getMessageCosts(messages))
}
