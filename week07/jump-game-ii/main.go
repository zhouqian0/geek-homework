package main

// 动态规划版本。
// 时间复杂度 O(n^2)，n 是 nums 长度。因为每遍历到一个元素时，都要再做一个 for 循环，循环这个元素能到的所有元素，并比较出一个最小的行进次数，所以复杂度是 O(n^2)。
// 空间复杂度 O(n)。因为需要一个 dp 数组保存递推过程中状态的变化情况，dp 数组长度 n，所以空间复杂度 O(n)。
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	// 因为要计算一个最小行进次数，所以默认赋予一个极大值
	for i := range dp {
		dp[i] = 1e9
	}
	// 初始化边界，因为第一次还没有走
	dp[0] = 0
	for i := range nums {
		// 取遍历 nums[i] 能到的所有元素，并比较达到他们的最小行进次数
		for j := i; j < min(i+nums[i]+1, n); j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 贪心版本。
// 相比之下动规版本，时间复杂度也在 O(n^2)，但是空间复杂度却是 O(1)，因为只用了常数级别的辅助变量。
func jump2(nums []int) (ans int) {
	n := len(nums)
	// 当前的位置下标
	now := 0
	// 只要还没到底
	for now < n - 1{
		// 获取一个当前能走到的最远下标
		right := now + nums[now]
		// 如果走完了数组，就更新步数并 return
		if right >= n -1 {
			ans++
			break
		}
		// 在 [now+1, right] 中找一个能前进最远的点，并保存到 next 中
		next, nextRange := now, right
		for i := now + 1; i <= right; i++ {
			if i + nums[i] > nextRange {
				nextRange = i + nums[i]
				next = i
			}
		}
		// 走到那个最远点
		now = next
		// 更新步数
		ans++
	}
	return
}