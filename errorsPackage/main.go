package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(10, 200)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
	return

}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}
