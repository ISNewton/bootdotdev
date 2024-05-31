package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func getSenderLog(s sender) string {
	return fmt.Sprintf(`
====================================
Sender name: %v
Sender number: %v
Sender rateLimit: %v
====================================
`, s.name, s.number, s.rateLimit)
}

func main() {
	output := getSenderLog(sender{rateLimit: 20, user: user{name: "John", number: 42}})
	fmt.Println(output)
}
