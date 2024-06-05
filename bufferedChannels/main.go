package main

func addEmailsToQueue(emails []string) chan string {
	myChannel := make(chan string, len(emails))

	for _, email := range emails {
		myChannel <- email
	}

	return myChannel
}
