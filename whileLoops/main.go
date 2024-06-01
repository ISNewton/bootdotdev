package main

import "fmt"

func main() {
	fmt.Println(getMaxMessagesToSend(1.35, 25))
}
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies < float64(maxCostInPennies) {

		maxMessagesToSend++
		actualCostInPennies *= costMultiplier

	}
	return maxMessagesToSend
}
