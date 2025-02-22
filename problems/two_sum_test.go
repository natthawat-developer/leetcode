
package problems

import (
    "testing"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        result []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},         // ตัวอย่างแรก: คำตอบคือ index 0 และ 1
        {[]int{3, 2, 4}, 6, []int{1, 2}},              // ตัวอย่างที่สอง: คำตอบคือ index 1 และ 2
        {[]int{3, 3}, 6, []int{0, 1}},                 // ตัวอย่างที่สาม: คำตอบคือ index 0 และ 1
        {[]int{1, 5, 6, 7}, 12, []int{1, 3}},          // ตัวอย่างที่สี่: คำตอบคือ index 1 และ 3
    }

    for _, tt := range tests {
        t.Run("Testing twoSum", func(t *testing.T) {
            got := twoSum(tt.nums, tt.target)
            if len(got) != len(tt.result) {
                t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.result)
                return
            }
            for i := range got {
                if got[i] != tt.result[i] {
                    t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.result)
                    return
                }
            }
        })
    }
}
