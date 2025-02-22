package problems

import "testing"

func TestReverse(t *testing.T) {
    tests := []struct {
        input  int
        output int
    }{
        {123, 321},          // 123 -> 321
        {-123, -321},        // -123 -> -321
        {120, 21},           // 120 -> 21
        {0, 0},              // 0 -> 0
        {1534236469, 0},     // Out of bounds, should return 0
    }

    for _, tt := range tests {
        t.Run("Testing reverse", func(t *testing.T) {
            got := reverse(tt.input)
            if got != tt.output {
                t.Errorf("reverse(%d) = %d; want %d", tt.input, got, tt.output)
            }
        })
    }
}
