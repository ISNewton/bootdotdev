package main

import "fmt"

type user struct {
	id    int
	name  string
	email string
}

type message struct {
	id       int
	content  string
	sender   user
	receiver user
}

func sendMessage(msg message) {
	fmt.Printf("Sender: %s\n", msg.sender.name)
	fmt.Printf("receiver: %s\n", msg.receiver.name)
	fmt.Printf("Sending message: %s\n", msg.content)
}
func main() {

	mySender := user{id: 1, name: "John Doe", email: "johndoe@gmail.com"}

	myReceiver := user{id: 2, name: "Ash", email: "ash@gmail.com"}

	myMessage := message{id: 22, content: "Hello World", sender: mySender, receiver: myReceiver}

	sendMessage(myMessage)

}
