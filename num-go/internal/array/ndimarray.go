package array

import (
	"fmt"
	"strings"

	"github.com/sukhman31/num-go/internal/utils"
)

type Array struct {
	data    []float64
	shape   []int
	strides []int
}

func NewArray(data []float64, shape []int) (*Array, error) {
	if len(data) != utils.Product(shape) {
		return nil, fmt.Errorf("data does not match the input shape")
	}

	strides := make([]int, len(shape))
	stride := 1
	for i := len(shape) - 1; i >= 0; i-- {
		strides[i] = stride
		stride *= shape[i]
	}

	return &Array{
		data:    data,
		shape:   shape,
		strides: strides,
	}, nil
}

func (a *Array) Set(data float64, indices ...int) error {
	index, err := getValidIndex(indices, *a)
	if err != nil {
		return err
	}
	a.data[index] = data
	return nil
}

func (a *Array) At(indices ...int) (float64, error) {
	index, err := getValidIndex(indices, *a)
	if err != nil {
		return 0, err
	}
	return a.data[index], nil
}

func (a *Array) Shape() []int {
	return a.shape
}

func (a *Array) PrettyPrint() string {
	return a.prettyPrintRecursive(0, []int{})
}

func (a *Array) prettyPrintRecursive(dimension int, indices []int) string {
	if dimension == len(a.shape) {
		index, _ := getValidIndex(indices, *a)
		return fmt.Sprintf("%.4f,", a.data[index])
	}

	var result strings.Builder
	result.WriteString("[")

	for i := 0; i < a.shape[dimension]; i++ {
		if i>0 {
			result.WriteString(" ")
		}
		newIndices := append(indices, i)
		result.WriteString(a.prettyPrintRecursive(dimension+1, newIndices))
	}
	result.WriteString("]")
	return result.String()
}

func getValidIndex(indices []int, a Array) (int, error) {
	if len(indices) != len(a.shape) {
		return -1, fmt.Errorf("invalid amount of indices provided")
	}

	index := 0
	for i, idx := range indices {
		if idx < 0 || idx >= a.shape[i] {
			return -1, fmt.Errorf("invalid index at position %d", i)
		}
		index += idx * a.strides[i]
	}
	return index, nil
}
