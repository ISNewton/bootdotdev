package main

import "fmt"

func getCounts(messagedUsers []string, validUsers map[string]int) {
	// ?
	for _, username := range messagedUsers {
		if _, ok := validUsers[username]; ok {
			validUsers[username]++
		}

	}
}

func main() {
	messagedUsers := []string{"cersei", "tyrion", "jaime", "tyrion", "tyrion"}
	validUsers := map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0}

	getCounts(messagedUsers, validUsers)

	fmt.Println(validUsers)

}
