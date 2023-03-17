package medium

//lint:ignore U1000 Ignore unused function temporarily for debugging
func nextPermutation(nums []int) {
	n := len(nums)
	var k, l int
	for k = n - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}
	if k < 0 {
		for j := 0; j < n/2; j++ {
			nums[j], nums[n-1-j] = nums[n-1-j], nums[j]
		}
	} else {
		for l = n - 1; l > k; l-- {
			if nums[k] < nums[l] {
				break
			}
		}

		nums[k], nums[l] = nums[l], nums[k]

		for j := k + 1; j < n/2; j++ {
			nums[j], nums[n-1-j] = nums[n-1-j], nums[j]
		}
	}
}
