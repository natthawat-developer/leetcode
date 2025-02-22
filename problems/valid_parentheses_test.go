package problems

import (
    "testing"
)

func TestIsValid(t *testing.T) {
    tests := []struct {
        input  string
        output bool
    }{
        {"()", true},                     // วงเล็บถูกต้อง
        {"()[]{}", true},                 // วงเล็บถูกต้องหลายชนิด
        {"(}", false},                    // วงเล็บไม่ตรงกัน
        {"([)]", false},                  // วงเล็บไม่ตรงกัน
        {"{[]}", true},                   // วงเล็บถูกต้อง
        {"", true},                       // กรณีไม่มีวงเล็บ
        {"(((((((((())))))))))", true},   // วงเล็บถูกต้องซ้ำๆ
        {"([", false},                    // ขาดวงเล็บปิด
    }

    for _, tt := range tests {
        t.Run("Testing isValid", func(t *testing.T) {
            got := isValid(tt.input)
            if got != tt.output {
                t.Errorf("isValid(%q) = %v; want %v", tt.input, got, tt.output)
            }
        })
    }
}
