package main

import (
	"errors"
	"fmt"
)

func main() {
	names := []string{"Ash", "Moh"}
	phones := []int{123}

	output, err := getUserMap(names, phones)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	output := make(map[string]user, len(names))
	for i, name := range names {

		output[name] = user{name: name, phoneNumber: phoneNumbers[i]}

	}
	return output, nil
}

type user struct {
	name        string
	phoneNumber int
}
