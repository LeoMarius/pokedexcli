package pokeapi

import (
	"testing"
	"time"
)

// func TestCleanInput(t *testing.T) {

// 	cases := []struct {
// 		input    string
// 		expected []string
// 	}{
// 		{
// 			input:    "  hello world  ",
// 			expected: []string{"hello", "world"},
// 		},
// 		{
// 			input:    " coucou comment  ca va  ",
// 			expected: []string{"coucou", "comment", "ca", "va"},
// 		},
// 	}

// 	for _, c := range cases {
// 		actual := cleanInput(c.input)
// 		if len(actual) != len(c.expected) {
// 			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
// 			continue
// 		}
// 		for i := range actual {
// 			word := actual[i]
// 			expectedWord := c.expected[i]
// 			if word != expectedWord {
// 				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
// 			}
// 		}
// 	}
// }

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	time.Sleep(waitTime)

}
