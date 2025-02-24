package problems

import (
    "testing"
)

func TestReverseArray(t *testing.T) {
    tests := []struct {
        input    []int
        expected []int
    }{
        {input: []int{1, 2, 3}, expected: []int{3, 2, 1}},
        {input: []int{5, 10, 15, 20}, expected: []int{20, 15, 10, 5}},
        {input: []int{7, 8, 9, 10, 11}, expected: []int{11, 10, 9, 8, 7}},
        {input: []int{}, expected: []int{}},
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            result := ReverseArray(tt.input)
            for i := range result {
                if result[i] != tt.expected[i] {
                    t.Errorf("Expected %v, but got %v", tt.expected, result)
                }
            }
        })
    }
}
