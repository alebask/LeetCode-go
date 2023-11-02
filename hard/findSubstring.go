package hard

func findSubstring(s string, words []string) []int {
	ht_limits := make(map[string]int)
	uwlimit := 0
	for _, word := range words {
		if _, ok := ht_limits[word]; !ok {
			uwlimit++
		}
		ht_limits[word]++
	}

	slen := len(s)
	wlen := len(words[0])
	wcount := len(words)

	indexes := make([]int, 0)
	for i := 0; i < wlen; i++ {

		ht_counts := make(map[string]int)
		left, right := i, i
		uwcount := 0
		for right <= slen-wlen {
			for uwcount < uwlimit && right <= slen-wlen {
				subs := s[right : right+wlen]
				if limit, ok := ht_limits[subs]; ok {
					ht_counts[subs]++
					if ht_counts[subs] == limit {
						uwcount++
					}
				}
				right += wlen
			}
			for uwlimit == uwcount && left < right {
				if right-left == wlen*wcount {
					indexes = append(indexes, left)
				}
				subs := s[left : left+wlen]
				if limit, ok := ht_limits[subs]; ok {
					ht_counts[subs]--
					if ht_counts[subs] < limit {
						uwcount--
					}
				}
				left += wlen
			}
		}
	}

	return indexes
}
