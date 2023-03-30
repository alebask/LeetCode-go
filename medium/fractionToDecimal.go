package medium

import "strconv"

//lint:ignore U1000 Ignore unused function temporarily for debugging
func fractionToDecimal(numerator int, denominator int) string {

	n := numerator / denominator
	r := numerator % denominator

	if r == 0 {
		return strconv.Itoa(n)
	}

	res := ""
	if n < 0 {
		res += "-"
	}
	res += strconv.Itoa(n) + "."

	ht := make(map[int]int)

	for r != 0 {
		r = r * 10
		n = r / denominator
		r = r % denominator

		if n < 0 {
			n = -n
		}

		res += strconv.Itoa(n)

		if index, ok := ht[r]; ok {
			return res[:index] + "(" + res[index:] + ")"
		} else {
			ht[r] = len(res)
		}
	}

	return res
}
