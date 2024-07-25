package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewArray(t *testing.T) {
	testCases := []struct {
		data []float64
		shape []int
		want *Array
		expectedErr bool
	}{
		{
			[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			[]int{2, 3, 2},
			&Array{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				shape: []int{2, 3, 2},
				strides: []int{6, 2, 1},
			},
			false,
		},
	}

	for _, tc := range testCases {
		got, err := NewArray(tc.data, tc.shape)
		if tc.expectedErr && err == nil {
			t.Errorf("unexpected error")
		}
		assert.Equal(t, *got, *tc.want, "they should be equal")
	}
}

func Test_At(t *testing.T) {
	testCases := []struct {
		array *Array
		position []int
		want float64
		expectedErr bool
	}{
		{
			&Array{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				shape: []int{2, 3, 2},
				strides: []int{6, 2, 1},
			},
			[]int{0,1,0},
			float64(3),
			false,
		},
	}

	for _, tc := range testCases {
		got, err := tc.array.At(tc.position...)
		if tc.expectedErr && err == nil {
			t.Errorf("unexpected error")
		}
		assert.Equal(t, got, tc.want, "they should be equal")
	}
}