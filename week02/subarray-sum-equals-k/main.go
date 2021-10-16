// 力扣 560 和为 K 的子数组（https://leetcode-cn.com/problems/subarray-sum-equals-k/）
package main

// 时间复杂度 O(n)，n 为 nums 长度，因为只需要一次遍历就可以统计出结果，每次遍历中对无序集合的查找耗时 O(1)，可以忽略，故而总耗时为 O(n)。
// 空间复杂度 O(n)，n 为 nums 长度，因为需要一个无序集合保存前缀和，而这个无序集合容量为 n + 1，故而总空间消耗为 O(n)。
func subarraySum(nums []int, k int) (ans int) {
	// 定义一个无序集合保存前缀和及其出现的次数
	hash := make(map[int]int)
	// 初始化无序集合，hash[0] 表示什么元素都没有添加时的前缀和计数
	hash[0] = 1
	// 初始化前缀和
	pre := 0
	// 遍历元素
	for _, v := range nums {
		// 计算前缀和
		pre += v
		// 查找在当前位置，是否存在满足条件的子数组
		if v, ok := hash[pre-k]; ok {
			// 更新计数
			ans += v
		}
		// 更新无序集合
		hash[pre]++
	}
	return
}
