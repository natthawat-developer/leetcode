package problems

import "testing"

func TestMaxProfit(t *testing.T) {
    tests := []struct {
        input    []int
        expected int
    }{
        {input: []int{7, 1, 5, 3, 6, 4}, expected: 7}, // ซื้อที่ 1 ขายที่ 5, ซื้อที่ 3 ขายที่ 6
        {input: []int{1, 2, 3, 4, 5}, expected: 4}, // ซื้อที่ 1 ขายที่ 5
        {input: []int{7, 6, 4, 3, 1}, expected: 0}, // ไม่มีกำไร เพราะราคาลดลงเรื่อยๆ
        {input: []int{2, 1, 2, 0, 1}, expected: 2}, // ซื้อที่ 1 ขายที่ 2, ซื้อที่ 0 ขายที่ 1
		{input: []int{1, 2, 4, 2, 5, 7}, expected: 8}, // ซื้อที่ 1 ขายที่ 2, ซื้อที่ 2 ขายที่ 4, ซื้อที่ 4 ขายที่ 5
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            result := MaxProfit(tt.input)
            if result != tt.expected {
                t.Errorf("Expected %d, but got %d", tt.expected, result)
            }
        })
    }
}
