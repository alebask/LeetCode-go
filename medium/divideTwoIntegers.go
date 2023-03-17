package medium

import "math"

// lint:ignore
func divide(dividend int, divisor int) int {

	res := 0

	n := abs(dividend)
	d := abs(divisor)

	negative := false
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		negative = true
	}

	for n >= d {
		i := 1
		for (n >> i) > d {
			i++
		}
		res += 1 << (i - 1)
		n -= d << (i - 1)
	}

	if negative {
		res = -res
	} else if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
