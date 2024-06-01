package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	if plan == planFree {
		return messages[1:3], nil
	}

	if plan == planPro {
		return messages[:], nil
	}

	return nil, errors.New("unsupported plan")
}

func main() {
	myMessages := [3]string{planFree, "World", "In Go"}
	messges, err := getMessageWithRetriesForPlan("My", myMessages)
	if err != nil {

		fmt.Println("Something went wrong: ", err)
		return
	}
	fmt.Println("Messages allowed: " + string(rune(len(messges))))
}
