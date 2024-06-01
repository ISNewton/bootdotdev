package main

import "fmt"

func main() {

	fmt.Println(getSMSErrorString(20, "Me"))
}

func getSMSErrorString(cost float64, recipient string) string {
	// ?
	return fmt.Sprintf("SMS that costs $%.1f to be sent to '%s' can not be sent\n", cost, recipient)
}
