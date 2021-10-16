// 力扣 697 数组的度（https://leetcode-cn.com/problems/degree-of-an-array/）
package main

type pos struct {
	count int
	start int
	end   int
}

// 时间复杂度：O(n)，n 为 nums 长度。执行了两次 for 循环，但它们并非嵌套执行，且他们的执行次数都是 n，故而时间复杂度为 O(n)。
// 空间复杂度：O(n)，n 为 nums 长度。因为需要一个额外的无序集合保存 nums 中出现的每个元素的相关信息，所以这个无序集合的长度不会超过 n，故空间复杂度 O(n)。
func findShortestSubArray(nums []int) (ans int) {
	hash := make(map[int]*pos)
	for i, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num].count++ // 更新当前元素出现的次数
			hash[num].end = i // 更新当前元素最后一次出现时的数组下标
			continue
		}
		hash[num] = &pos{
			count: 1, // 当前元素出现次数
			start: i, // 当前元素第一次出现时的数组下标
			end:   i, // 当前元素最后一次出现时的数组下标
		}
	}

	max := 0    // 数组的度
	ans = 49999 // nums[i] 是一个在 0 到 49,999 范围内的整数。
	for _, num := range nums {
		if max < hash[num].count { // 当前的元素出现的次数大于之前每一个元素的出现次数
			max = hash[num].count                     // 更新数组的度
			ans = hash[num].end - hash[num].start + 1 // 更新结果
		} else if max == hash[num].count { // 如果当前元素的出现的次数和当前统计出的数组的度相同
			ans = min(ans, hash[num].end-hash[num].start+1) // 更新结果（因为需要查找最短连续子数组）
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
