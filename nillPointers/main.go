package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return

	}
	fmt.Println(message)
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

func main() {
	var message *string

	removeProfanity(message)
}
