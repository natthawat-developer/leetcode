package problems

import "testing"

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
    tests := []struct {
        input    []int
        expected []int
        length   int
    }{
        {input: []int{1, 1, 2}, expected: []int{1, 2}, length: 2},
        {input: []int{0, 0, 1, 1, 1, 2, 2, 3}, expected: []int{0, 1, 2, 3}, length: 4},
        {input: []int{1, 1, 1, 1}, expected: []int{1}, length: 1},
        {input: []int{}, expected: []int{}, length: 0},
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            length := RemoveDuplicatesFromSortedArray(tt.input)
            if length != tt.length {
                t.Errorf("Expected length %d, but got %d", tt.length, length)
            }

            for i := 0; i < length; i++ {
                if tt.input[i] != tt.expected[i] {
                    t.Errorf("Expected %v, but got %v", tt.expected, tt.input)
                    break
                }
            }
        })
    }
}
