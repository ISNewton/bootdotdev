package main

func concurrentFib(n int) []int {
	// ?
	myChannel := make(chan int)

	mySlice := []int{}

	go fibonacci(n, myChannel)

	for value := range myChannel {
		mySlice = append(mySlice, value)
	}

	return mySlice
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
