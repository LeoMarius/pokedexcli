package pokecache

import (
	"fmt"
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

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	time.Sleep(waitTime)
}
