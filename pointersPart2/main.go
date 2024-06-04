package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
}

func main() {
	myString := "Oh man I've seen some crazy ass shiz in my time..."
	removeProfanity(&myString)
	fmt.Println(myString)
}
