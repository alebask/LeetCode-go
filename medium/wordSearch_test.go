package medium

import (
	"fmt"
	"testing"
)

func Test_wordSearch(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCCED",
			expected: true},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%+v %+v --> %v", test.board, test.word, test.expected)

		t.Run(testName, func(t *testing.T) {
			actual := wordSearchChannels(test.board, test.word)
			if actual != test.expected {
				t.Errorf("expected %v, actual %v", test.expected, actual)
			}
		})
	}
}
