package problems

import "testing"

func TestCountPrimes(t *testing.T) {
    tests := []struct {
        input    int
        expected int
    }{
        {input: 10, expected: 4},  // 2, 3, 5, 7
        {input: 0, expected: 0},   // ไม่มีจำนวนเฉพาะ
        {input: 1, expected: 0},   // ไม่มีจำนวนเฉพาะ
        {input: 2, expected: 0},   // ไม่มีจำนวนเฉพาะที่น้อยกว่า 2
        {input: 20, expected: 8},  // 2, 3, 5, 7, 11, 13, 17, 19
        {input: 100, expected: 25}, // จำนวนเฉพาะที่ < 100
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            result := CountPrimes(tt.input)
            if result != tt.expected {
                t.Errorf("CountPrimes(%d) = %d; expected %d", tt.input, result, tt.expected)
            }
        })
    }
}
