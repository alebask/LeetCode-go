package medium

import "sort"

//lint:ignore U1000 Ignore unused function temporarily for debugging
func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	res := 1<<31 - 1
	mindiff := res

	for i, n := range nums {

		start := i + 1
		end := len(nums) - 1

		for start < end {
			sum := n + nums[start] + nums[end]

			if sum == target {
				return target
			} else if sum < target {
				if target-sum < mindiff {
					mindiff = target - sum
					res = sum
				}
				start++
			} else {
				if sum-target < mindiff {
					mindiff = sum - target
					res = sum
				}
				end--
			}
		}
	}

	return res

}
