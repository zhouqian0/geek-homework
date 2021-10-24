package main

import "sort"

// 时间复杂度 O(n*(n!))，n 为数组长度，因为总共有 n！个全排列可能，且需要 n 的时间复制结果数组。
// 空间复杂度 O(n)，需要两个额外的长度 n 的数组辅助计算，且递归深度为 n，所以总空间复杂度为 O(n)。
func permuteUnique(nums []int) (ans [][]int) {
	// ans 长度不超过 len(nums) 的阶乘
	// 做一次快排，让所有元素有序，使相同的元素相邻，方便筛选
	sort.Ints(nums)

	n := len(nums)
	// 存放排列结果的公共变量
	// 使用公共变量需要注意还原现场
	set := make([]int, 0, n)
	used := make([]bool, n)

	var dfs func(pos int)
	dfs = func(pos int) {
		// 递归终止条件，即已经加入了所有元素
		if pos == n {
			// 更新答案
			ans = append(ans, append([]int{}, set...))
			return
		}
		for i := 0; i < n; i++ {
			// 如果当前元素已经加入排列，或者当前元素和前一个元素相同，但前一个元素未加入排列
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			// 记录当前元素已经加入排列
			used[i] = true
			// 将当前元素加入排列
			set = append(set, nums[i])
			// 计算结果
			dfs(pos + 1)
			// 还原现场
			used[i] = false
			set = set[:len(set)-1]
		}
	}
	// 开是递归
	dfs(0)
	return
}
