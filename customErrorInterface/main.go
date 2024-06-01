package main

import "fmt"

func main() {
	result, err := divide(18, 0)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

type divideError struct {
	dividend float64
}

func (theError divideError) Error() string {
	return fmt.Sprintf("can not divide %f by zero\n", theError.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
