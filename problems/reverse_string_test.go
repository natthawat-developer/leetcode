package problems

import "testing"

func TestReverseString(t *testing.T) {
    // Test case 1: Reverse string "hello"
    input := []byte("hello")
    expected := "olleh"
    reverseString(input)
    if string(input) != expected {
        t.Errorf("Expected %v, but got %v", expected, string(input))
    }

    // Test case 2: Reverse string "abc"
    input = []byte("abc")
    expected = "cba"
    reverseString(input)
    if string(input) != expected {
        t.Errorf("Expected %v, but got %v", expected, string(input))
    }

    // Test case 3: Reverse string "a"
    input = []byte("a")
    expected = "a"
    reverseString(input)
    if string(input) != expected {
        t.Errorf("Expected %v, but got %v", expected, string(input))
    }

    // Test case 4: Reverse empty string
    input = []byte("")
    expected = ""
    reverseString(input)
    if string(input) != expected {
        t.Errorf("Expected %v, but got %v", expected, string(input))
    }
}
