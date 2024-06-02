package main

import (
	"fmt"
	"strings"
)

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?
	for _, message := range msg {
		for b, badWord := range badWords {

			if strings.Contains(message, badWord) {
				return b
			}
		}

	}
	return -1
}

func main() {
	badWords := []string{"hello", "world", "This", "Go"}
	msg := []string{"hello", "ssh"}

	fmt.Println(indexOfFirstBadWord(msg, badWords))

}
