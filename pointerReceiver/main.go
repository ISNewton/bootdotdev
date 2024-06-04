package main

import "fmt"

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func main() {
	myEmail := email{message: "Hello World"}
	fmt.Println(myEmail.message)

	myEmail.setMessage("New hello")

	fmt.Println(myEmail.message)
}
