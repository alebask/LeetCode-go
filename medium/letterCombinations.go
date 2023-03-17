package medium

//lint:ignore U1000 Ignore unused function temporarily for debugging
func letterCombinations1(digits string) []string {

	n := len(digits)

	if n == 0 {
		return []string{}
	}

	buttons := []string{"",
		"", "abc", "def",
		"ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz"}

	combine := func(seed []string, letters string) []string {
		res := make([]string, 0)
		for _, s := range seed {
			for _, l := range letters {
				res = append(res, s+string(l))
			}
		}
		return res
	}

	res := []string{""}
	for _, d := range digits {
		letters := buttons[d-'0']
		res = combine(res, letters)
	}
	return res
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func letterCombinations2(digits string) []string {

	n := len(digits)

	if n == 0 {
		return []string{}
	}

	buttons := []string{"",
		"", "abc", "def",
		"ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz"}

	combine := func(seed []string, letters string) []string {
		res := make([]string, 0)
		for _, s := range seed {
			for _, l := range letters {
				res = append(res, s+string(l))
			}
		}
		return res
	}

	res := []string{""}
	for _, d := range digits {
		letters := buttons[d-'0']
		res = combine(res, letters)
	}
	return res
}
