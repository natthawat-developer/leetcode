package problems

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			input: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			original := make([][]int, len(tt.input))
			for i := range tt.input {
				original[i] = append([]int(nil), tt.input[i]...) // Clone the original input
			}

			RotateImage(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("For input %v, expected %v but got %v", original, tt.expected, tt.input)
			}
		})
	}
}
