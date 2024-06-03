package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	outputMap := make(map[rune]map[string]int)

	for _, name := range names {
		if _, ok := outputMap[rune(name[0])]; ok {
			runeMap := outputMap[rune(name[0])]
			runeMap[name]++
		} else {
			runeMap := make(map[string]int)
			runeMap[name] = 1
			outputMap[rune(name[0])] = runeMap

		}
	}
	return outputMap
}

func main() {
	theNames := []string{
		"billy",
		"billy",
		"bob",
		"joe",
	}

	fmt.Println(getNameCounts(theNames))

}
