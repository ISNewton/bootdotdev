package main

import (
	"fmt"
)

func main() {

	cost, err := sendSMSToCouple("Messages sent successfully at the cost : ", "Hello world")
	if err != nil {
		fmt.Println("Something went wrong: " + err.Error())
	} else {

		fmt.Println("Messages sent successfully at the cost : ", cost)
	}
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	// ?
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	secondCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	return cost + secondCost, nil

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
