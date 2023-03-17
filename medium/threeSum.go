package medium

import "sort"

//lint:ignore U1000 Ignore unused function temporarily for debugging
func threeSum1(nums []int) [][]int {

	type Key struct {
		x int
		y int
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	triplets := make(map[Key]bool, 0)

	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		m := make(map[int]bool, 0)
		for j := i + 1; j < len(nums); j++ {
			n2 := nums[j]
			n3 := -n1 - n2
			if exists := m[n3]; exists {
				x := min(n1, min(n2, n3))
				y := max(n1, max(n2, n3))
				triplets[Key{x: x, y: y}] = true
			} else {
				m[n2] = true
			}
		}
	}

	res := make([][]int, len(triplets))
	i := 0
	for k := range triplets {
		res[i] = []int{k.x, k.y, -k.x - k.y}
		i++
	}
	return res
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func threeSum2(nums []int) [][]int {

	sort.Ints(nums)
	res := make([][]int, 0)

	for i, n := range nums {

		if i > 0 && n == nums[i-1] {
			continue
		}

		start := i + 1
		end := len(nums) - 1

		for start < end {
			sum := n + nums[start] + nums[end]
			if sum == 0 {
				res = append(res, []int{n, nums[start], nums[end]})
				start++
				end--
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}
			} else if sum < 0 {
				start++
				for start < end && nums[start] == nums[start+1] {
					start++
				}
			} else {
				end--
				for start < end && nums[end] == nums[end-1] {
					end--
				}
			}
		}

	}

	return res
}
