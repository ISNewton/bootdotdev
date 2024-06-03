package main

import (
	"errors"
	"fmt"
)

func main() {

	users := map[string]user{
		"ash": {name: "ash", number: 123, scheduledForDeletion: true},
	}

	isDeleted, err := deleteIfNecessary(users, "ash")
	fmt.Println(isDeleted, err)

}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	theName, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !theName.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)

	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
