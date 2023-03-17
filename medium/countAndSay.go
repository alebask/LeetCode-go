package medium

import (
	"strconv"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func countAndSay(n int) string {

	seed := "1"
	newSeed := ""

	for i := 1; i < n; i++ {
		for j := 0; j < len(seed); j++ {
			count := 1
			for j < len(seed)-1 && seed[j] == seed[j+1] {
				count++
				j++
			}
			newSeed += strconv.Itoa(count) + strconv.Itoa(int(seed[j]))
		}

		seed = newSeed
		newSeed = ""
	}

	return seed
}
