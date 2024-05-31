package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (ai authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + ai.username + ":" + ai.password
}
func main() {

	user := authenticationInfo{username: "ash", password: "my1212"}

	fmt.Println(user.getBasicAuth())
}
