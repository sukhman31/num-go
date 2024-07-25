package array

import (
	"fmt"
	"math"
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

	return &Array{
		data:    data,
		shape:   shape,
		strides: getStrides(shape),
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

func Zeros(shape []int) *Array{
	return &Array{
		data: make([]float64,utils.Product(shape)),
		shape: shape,
		strides: getStrides(shape),
	}
}

func Ones(shape []int) *Array {
	data := make([]float64,utils.Product(shape))
	for idx := range data {
		data[idx] = 1
	}
	return &Array{
		data: data,
		shape: shape,
		strides: getStrides(shape),
	}
}

func Arange(args ...interface{}) (*Array, error) {
	switch (len(args)) {
	case 1:
		maxVal, ok := args[0].(int)
		if !ok {
			return nil, fmt.Errorf("value must be an integer")
		}
		return arangeOneArg(maxVal)
	case 2:
		minVal, ok1 := args[0].(int)
		maxVal, ok2 := args[1].(int)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("both values must be integers")
		}
		return arangeTwoArgs(minVal, maxVal)
	case 3:
		minVal, ok1 := args[0].(int)
		maxVal, ok2 := args[1].(int)
		step, ok3 := args[2].(int)
		if !ok1 || !ok2 || !ok3 {
			return nil, fmt.Errorf("all values must be integers")
		}
		return arangeThreeArgs(minVal, maxVal, step)
	default:
        return nil, fmt.Errorf("arange accepts 1, 2, or 3 arguments")
	}
}


func arangeOneArg(maxVal int) (*Array, error) {
	return arange(0, maxVal, 1)
}

func arangeTwoArgs(minVal int, maxVal int) (*Array, error) {
	return arange(minVal, maxVal, 1)
}

func arangeThreeArgs(minVal int, maxVal int, step int) (*Array, error) {
	return arange(minVal, maxVal, step)
}

func arange(minVal int, maxVal int, step int) (*Array, error) {
	if minVal > maxVal {
		return nil, fmt.Errorf("minimum val cannot be greater than maximum value")
	}
	if step == 0 {
		return nil, fmt.Errorf("step cannot be 0")
	}
	arraySize := math.Ceil(float64((maxVal-minVal))/float64(step))
	data := make([]float64, int(arraySize))
	for i := range data {
		data[i] = float64(minVal + i*step)
	}
	return &Array{
		data: data,
		shape: []int{len(data), 1},
		strides: getStrides([]int{len(data), 1}),
	}, nil
}

func (a *Array) Shape() []int {
	return a.shape
}

func (a *Array) Ndim() int {
	return len(a.shape)
}

func (a *Array) Size() int {
	return utils.Product(a.shape)
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
			return -1, fmt.Errorf("invalid index at position: %d", i)
		}
		index += idx * a.strides[i]
	}
	return index, nil
}

func getStrides(shape []int) []int{
	strides := make([]int, len(shape))
	stride := 1
	for i := len(shape) - 1; i >= 0; i-- {
		strides[i] = stride
		stride *= shape[i]
	}

	return strides
}