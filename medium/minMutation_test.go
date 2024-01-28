package medium

import (
	"fmt"
	"testing"
)

func Test_minMutation(t *testing.T) {

	tests := []struct {
		startGene string
		endGene   string
		bank      []string
		expected  int
	}{

		{startGene: "AACCGGTT", endGene: "AACCGGTA", bank: []string{"AACCGGTA"}, expected: 1},
		{startGene: "AACCGGTT", endGene: "AAACGGTA", bank: []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, expected: 2},
		{startGene: "AACCGGTT", endGene: "AAACGGTA", bank: []string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"}, expected: 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s, %v", tt.startGene, tt.endGene, tt.bank)
		t.Run(testname, func(t *testing.T) {
			actual := minMutation(tt.startGene, tt.endGene, tt.bank)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}

}
