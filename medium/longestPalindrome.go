package medium

//lint:ignore U1000 Ignore unused function temporarily for debugging
func longestPalindrome(s string) string {

	len := len(s)
	if len == 1 {
		return s
	}
	if len == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}

	left, right := 0, 0
	prev, last := len-2, len-1

	processPalindrome := func(l, r int) {
		for l > 0 && r < last {
			if s[l-1] != s[r+1] {
				break
			}
			l--
			r++
		}
		if r-l > right-left {
			right = r
			left = l
		}
	}

	for i := 0; i < prev; i++ {
		//even
		if s[i] == s[i+1] {
			processPalindrome(i, i+1)
		}
		//odd
		if s[i] == s[i+2] {
			processPalindrome(i, i+2)
		}
	}

	if s[prev] == s[last] {
		processPalindrome(prev, last)
	}
	return s[left : right+1]
}
