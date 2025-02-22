package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input  int
        output bool
    }{
        {121, true},         // 121 is a palindrome
        {-121, false},       // -121 is not a palindrome
        {10, false},         // 10 is not a palindrome
        {12321, true},       // 12321 is a palindrome
        {0, true},           // 0 is a palindrome
    }

    for _, tt := range tests {
        t.Run("Testing isPalindrome", func(t *testing.T) {
            got := isPalindrome(tt.input)
            if got != tt.output {
                t.Errorf("isPalindrome(%d) = %v; want %v", tt.input, got, tt.output)
            }
        })
    }
}
