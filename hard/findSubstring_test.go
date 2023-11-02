package hard

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_findSubstring(t *testing.T) {
	tests := []struct {
		s        string
		words    []string
		expected []int
	}{
		{s: "barfoothefoobarman", words: []string{"foo", "bar"}, expected: []int{0, 9}},
		{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}, expected: []int{}},
		{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}, expected: []int{6, 9, 12}},
		{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "good"}, expected: []int{8}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s --> %v", tt.s, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := findSubstring(tt.s, tt.words)
			if !slices.Equal(actual, tt.expected) {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
