package medium

func minReorder(n int, connections [][]int) int {

	cities := make(map[int]city)

	for i := 0; i < n; i++ {
		cities[i] = city{val: i, prev: make(map[int]struct{}), next: make(map[int]struct{})}
	}

	for _, conn := range connections {
		from := conn[0]
		to := conn[1]
		cities[from].next[to] = struct{}{}
		cities[to].prev[from] = struct{}{}
	}

	res := 0

	for len(cities) > 1 {

		for k, city := range cities {
			if k == 0 {
				continue
			}

			if len(city.prev) == 0 && len(city.next) == 1 {
				for nk, _ := range city.next {
					delete(cities[nk].prev, k)
				}
				delete(cities, k)
			}
			if len(city.next) == 0 && len(city.prev) == 1 {
				for pk, _ := range city.prev {
					delete(cities[pk].next, k)
					res++
				}
				delete(cities, k)
			}
		}
	}

	return res
}

type city struct {
	val  int
	prev map[int]struct{}
	next map[int]struct{}
}
