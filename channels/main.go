package main

import (
	"fmt"
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)
	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	fmt.Println("h")
	return isOld
}

func main() {
	emails := [3]email{
		{body: "Hello World", date: time.Now()},
		{body: "Hi", date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC)},
		{body: "Shi", date: time.Now()},
	}
	outpu := checkEmailAge(emails)
	fmt.Println("outpu:", outpu)
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
