package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	var myCosts []float64

	for _, c := range costs {
		myCosts[c.day] = myCosts[c.day] + c.value
	}
	return myCosts
}

func main() {
	theCosts := []cost{
		{0, 4.0},
		{1, 2.1},
		{5, 2.5},
		{1, 3.1},
	}
	fmt.Println(getCostsByDay(theCosts))
}
