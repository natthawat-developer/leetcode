package problems

import (
    "testing"
)

func TestLongestCommonPrefix(t *testing.T) {
    tests := []struct {
        input  []string
        output string
    }{
        {[]string{"flower", "flow", "flight"}, "fl"},        // "flower", "flow", "flight" -> "fl"
        {[]string{"dog", "racecar", "car"}, ""},             // "dog", "racecar", "car" -> ""
        {[]string{"apple", "ape", "april"}, "ap"},           // "apple", "ape", "april" -> "ap"
        {[]string{"car", "cat", "cap"}, "ca"},               // "car", "cat", "cap" -> "ca"
        {[]string{"", "dog", "cat"}, ""},                    // "", "dog", "cat" -> ""
    }

    for _, tt := range tests {
        t.Run("Testing longestCommonPrefix", func(t *testing.T) {
            got := longestCommonPrefix(tt.input)
            if got != tt.output {
                t.Errorf("longestCommonPrefix(%v) = %q; want %q", tt.input, got, tt.output)
            }
        })
    }
}
