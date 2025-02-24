package problems

import (
    "testing"
    "reflect"
)

func TestFizzBuzz(t *testing.T) {
    tests := []struct {
        input    int
        expected []string
    }{
        {input: 5, expected: []string{"1", "2", "Fizz", "4", "Buzz"}},
        {input: 10, expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}},
        {input: 15, expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
        {input: 1, expected: []string{"1"}},
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            result := FizzBuzz(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("Expected %v, but got %v", tt.expected, result)
            }
        })
    }
}
