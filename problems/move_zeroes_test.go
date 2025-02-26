package problems

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
		{input: []int{0, 0, 1}, expected: []int{1, 0, 0}},
		{input: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}}, // ไม่มี 0
		{input: []int{0, 0, 0}, expected: []int{0, 0, 0}}, // มีแต่ 0
		{input: []int{0}, expected: []int{0}}, // กรณี 1 ตัว
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			nums := append([]int(nil), tt.input...) // Copy array เพื่อไม่ให้แก้ input โดยตรง
			MoveZeroes(nums)
			if !reflect.DeepEqual(nums, tt.expected) {
				t.Errorf("For input %v, expected %v but got %v", tt.input, tt.expected, nums)
			}
		})
	}
}
