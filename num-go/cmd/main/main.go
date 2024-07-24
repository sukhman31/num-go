package main

import (
	"fmt"

	"github.com/sukhman31/num-go/internal/array"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	shape := []int{2, 3, 2}

	arr, err := array.NewArray(data, shape)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print original value
	value, err := arr.At(1, 2, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Original value at [1,2,1]:", value)

	// Set new value
	err = arr.Set(100, 1, 2, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print new value
	value, err = arr.At(1, 2, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New value at [1,2,1]:", value)
	fmt.Println(arr.Shape())
}
