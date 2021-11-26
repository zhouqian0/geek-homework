package main

// 时间复杂度 O(n)，n 为数组 nums 长度，因为代码中最耗时的过程时遍历整个数组，所以时间复杂度为 O(n)。
// 空间复杂度 O(1)，因为只借助了变量 pre 辅助计算，所以空间复杂度为 O(1)。
func canJump(nums []int) bool {
	n := len(nums)
	// 用 farthest 来表示 在 nums[i] 上能到的最远范围
	// 因为每一步的能走的最远范围只和上一步有关，所以没有用数组，只用一个变量辅助计算
	// 状态转移方程为 farthest[i] = max(farthest-1, nums[i])
	// 边界为初始 farthest[0] = nums[0]，如果中间有任何一步 farthest[i] == 0 了，就说明无法走到终点
	// 目标为 farthest[len(nums) - 1]
	farthest := nums[0]
	for i := 1; i < n; i++ {
		if farthest == 0 {
			return false
		}
		farthest = max(farthest-1, nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
