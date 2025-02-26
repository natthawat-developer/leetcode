package problems

import "testing"

func TestSingleNumber(t *testing.T) {
    tests := []struct {
        input    []int
        expected int
    }{
        {input: []int{2, 2, 1}, expected: 1},
        {input: []int{4, 1, 2, 1, 2}, expected: 4},
        {input: []int{1}, expected: 1}, // กรณีมีตัวเดียว
        {input: []int{0, 0, 3}, expected: 3},
        {input: []int{7, 3, 5, 3, 7}, expected: 5},
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            result := SingleNumber(tt.input)
            if result != tt.expected {
                t.Errorf("For input %v, expected %d but got %d", tt.input, tt.expected, result)
            }
        })
    }
}
