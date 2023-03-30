package medium

//lint:ignore U1000 Ignore unused function temporarily for debugging
func palondromePartitioning(s string) [][]string {

	res := make([][]string, 0)
	res = append(res, []string{})

	for _, rn := range s {

		char := string(rn)
		lr := len(res)
		for i := 0; i < lr; i++ {

			item := res[i]
			last := len(item) - 1
			prev := last - 1

			if last >= 0 && item[last] == char {
				p1 := make([]string, len(item))
				copy(p1, item)
				p1[last] = item[last] + char
				res = append(res, p1)
			}
			if last >= 1 && item[prev] == char {
				p2 := make([]string, len(item)-1)
				copy(p2, item)
				p2[prev] = item[prev] + item[last] + char
				res = append(res, p2)
			}

			res[i] = append(item, char)
		}
	}

	return res
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func palondromePartitioningChannels(s string) [][]string {

	toCharStream := func(s string) <-chan string {
		out := make(chan string)
		go func() {
			for _, r := range s {
				out <- string(r)
			}
			close(out)
		}()
		return out
	}

	appendChar := func(char string, in <-chan []string) <-chan []string {
		out := make(chan []string)
		go func() {
			defer close(out)
			for p := range in {
				last := len(p) - 1
				prev := last - 1

				out <- append(p, char)

				if last >= 0 && p[last] == char {
					p1 := make([]string, len(p))
					copy(p1, p)
					p1[last] = p[last] + char
					out <- p1
				}
				if last >= 1 && p[prev] == char {
					p2 := make([]string, len(p)-1)
					copy(p2, p)
					p2[prev] = p[prev] + p[last] + char
					out <- p2
				}
			}
		}()
		return out
	}

	emptySeed := func() <-chan []string {
		out := make(chan []string)
		go func() {
			defer close(out)
			out <- []string{}
		}()
		return out
	}

	stream := emptySeed()

	for char := range toCharStream(s) {
		stream = appendChar(char, stream)
	}

	res := make([][]string, 0)
	for r := range stream {
		res = append(res, r)
	}

	return res
}
