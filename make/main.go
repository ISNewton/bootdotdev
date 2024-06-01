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
	messages := []string{"Lorem skdjfksjkfjsdfjkjkskdjfksdf ksjdfksjfksjdfk skdfjksdjfksdjfkjsd fkjskdfjsdkfjskdfjsdk f hekrksjfksjfksjfksjfskjfkajdfk",
		"world"}

	fmt.Println(getMessageCosts(messages))
}
