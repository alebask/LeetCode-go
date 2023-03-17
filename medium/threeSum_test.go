package medium

import "testing"

func Benchmark_threeSum(b *testing.B) {

	for i := 0; i < b.N; i++ {
		threeSum1([]int{-1, 0, 1, 2, -1, -4})
	}
}
