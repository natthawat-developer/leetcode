package problems

import (
    "reflect"
    "testing"
)

func TestRotateArray(t *testing.T) {
    tests := []struct {
        input    []int
        k        int
        expected []int
    }{
        {input: []int{1, 2, 3, 4, 5, 6}, k: 3, expected: []int{4, 5, 6, 1, 2, 3}},
        {input: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{5, 6, 7, 1, 2, 3, 4}},
        {input: []int{-1, -100, 3, 99}, k: 2, expected: []int{3, 99, -1, -100}},
        {input: []int{1, 2}, k: 1, expected: []int{2, 1}},  // Small array case
        {input: []int{1, 2, 3}, k: 0, expected: []int{1, 2, 3}},  // No rotation
        {input: []int{1, 2, 3, 4}, k: 4, expected: []int{1, 2, 3, 4}},  // Full rotation
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            nums := append([]int(nil), tt.input...) // Create a copy to avoid modifying original test case
            RotateArray(nums, tt.k)
            if !reflect.DeepEqual(nums, tt.expected) {
                t.Errorf("For input %v with k=%d, expected %v but got %v", tt.input, tt.k, tt.expected, nums)
            }
        })
    }
}
