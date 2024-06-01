package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// ?
	messages := [3]string{primary, secondary, tertiary}

	costs := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}

	return messages, costs
}

func main() {
	fmt.Println(getMessageWithRetries("hello", "world", "!"))
}
