package triangle

// 时间复杂度 O(n^2)，n 为 triangle 的长度，因为需要遍历每一层的每一个元素来计算最小路径，所以 时间复杂度在 O(n^n)。
// 空间复杂度 O(n), 因为需要一个数组存储计算的过程，这个数组的长度为 n，故而空间复杂度在 O(n)。
func minimumTotal(triangle [][]int) int {
	// 获取长度
	n := len(triangle)
	// 直接用数组的最后一层用作辅助动规
	dp := triangle[n-1]

	// 这里采用逆思维，用数组的最后一层做基准，直接从倒数第二层开始遍历
	for i := n - 2; i >= 0; i-- {
		// 遍历当前层所有元素
		for j, v := range triangle[i] {
			// 对比当当前元素和下一层对应元素和下一层后一个元素那个更小
			// 因为路径的移动只能从前行的下标 i ，移动到下一行的下标 i 或 i + 1
			// 所以取 dp[j], dp[j + 1] 中相对小的那个（dp 数组存储得失下一层到数组末层的最小路径）
			dp[j] = min(dp[j], dp[j+1]) + v
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
