package problems

import (
    "reflect"
    "testing"
)

func TestIntersect(t *testing.T) {
    tests := []struct {
        nums1    []int
        nums2    []int
        expected []int
    }{
        {nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, expected: []int{2, 2}},
        {nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, expected: []int{4, 9}},
        {nums1: []int{1, 2, 3}, nums2: []int{4, 5, 6}, expected: []int{}}, // ไม่มีตัวที่ตรงกัน
        {nums1: []int{1, 1, 1, 2}, nums2: []int{1, 1}, expected: []int{1, 1}}, // ซ้ำหลายตัว
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            result := Intersect(tt.nums1, tt.nums2)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("For nums1=%v and nums2=%v, expected %v but got %v", tt.nums1, tt.nums2, tt.expected, result)
            }
        })
    }
}
