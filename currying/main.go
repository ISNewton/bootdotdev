package main

import "fmt"

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(x string, y string) {
		fmt.Println(formatter(x, y))
	}
}
