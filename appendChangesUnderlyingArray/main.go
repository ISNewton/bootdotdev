package main

import "fmt"

func main() {
	myFirstSlice := make([]int, 2, 6)
	myFirstSlice[0] = 4
	myFirstSlice[1] = 7
	fmt.Println(myFirstSlice)

	mySecondSlice := append(myFirstSlice, 10)
	myThirdSlice := append(myFirstSlice, 11)

	fmt.Println(myFirstSlice)
	fmt.Println(mySecondSlice)
	fmt.Println(myThirdSlice)
	// secondSlice 10 value on the index 2 is has been overwritten by Third slice 11 value
	// Now secondSlice[2] holds the value: 11 instead of 10
	// instead it is best practice to hold the slice returned by append() in the same slice as:
	myFirstSlice = append(myFirstSlice, 66)
}
