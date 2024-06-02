package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// ?
	myMatrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			myMatrix[i] = append(myMatrix[i], j*i)
		}

	}

	return myMatrix
}
func main() {
	fmt.Println(createMatrix(10, 10))
}
