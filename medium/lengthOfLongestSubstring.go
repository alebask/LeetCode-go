package medium

//lint:ignore U1000 Ignore unused function temporarily for debugging
func lengthOfLongestSubstring(s string) int {

	if len(s) < 2 {
		return len(s)
	}

	lastSeen := make([]int, 128)
	for i := range lastSeen {
		lastSeen[i] = -1
	}

	lmax := 1
	for start, end := 0, 0; end < len(s); end++ {
		i := s[end]
		if lastSeen[i] != -1 {
			if start <= lastSeen[i] {
				start = lastSeen[i] + 1
			}
		}
		lastSeen[i] = end

		len := end - start + 1
		if len > lmax {
			lmax = len
		}
	}

	return lmax
}
