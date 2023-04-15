package medium

func lengthOfLIS(nums []int) int {

	res := 0
	for i := 0; i < len(nums); i++ {
		if res == 0 || nums[res-1] < nums[i] {
			nums[res] = nums[i]
			res++
		} else {
			for j := 0; j < res; j++ {
				if nums[j] >= nums[i] {
					nums[j] = nums[i]
					break
				}
			}
		}
	}

	return res
}
