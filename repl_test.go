package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "The Priory of the Orange Tree",
			expected: []string{"the", "priory", "of", "the", "orange", "tree"},
		},
		{
			input:    "    ThiS  hAs   manY         SPaces   ",
			expected: []string{"this", "has", "many", "spaces"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Result []string length does not match expected")
			fmt.Println("Test failed")
		}
		for i := range actual {
			word := actual[i]
			expectedword := c.expected[i]
			if word != expectedword {
				t.Errorf("Word in result does not match expected: %v != %v", word, expectedword)
				fmt.Println("Test failed")
			}
		}
	}
}
