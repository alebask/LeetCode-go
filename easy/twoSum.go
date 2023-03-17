package easy

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	res := make([]int, 2)

	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			res[0] = j
			res[1] = i
			break
		}
		m[v] = i
	}

	return res
}
