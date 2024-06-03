package main

import "fmt"

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type User struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]User, name string) (log string) {
	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	defer delete(users, name)
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

func main() {
	users := map[string]User{"ash": User{"ash", 100, false}}
	fmt.Println(users)

	logAndDelete(users, "ash")

	fmt.Println(users)

}
