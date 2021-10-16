// 力扣 811 子域名访问计数（https://leetcode-cn.com/problems/subdomain-visit-count/）
package main

// 用于存放数组中每个元素的统计信息
type element struct {
	start int // 元素第一次出现的下标
	end   int // 元素最后一次出现的下标
	count int // 元素出现的次数
}

// 时间复杂度：O(n)。n 为 nums 长度，因为经历了两次 for 循环，但这两次 for 循环顺序执行而非嵌套，第一次 for 循环耗时 n， 第二次 for 循环 不超过 n，故而总时间复杂度为 O(n)。
// 空间复杂度：O(n)。n 为 nums 长度，因为需要一个无序集合存储所有 num 的相关信息结构体，而这些结构体数量不超过 n，故而空间复杂度为 O(n)。
func findShortestSubArray(nums []int) int {
	maxCount := 1
	// 定义一个无序集合存放元素的值及其相关统计信息
	hash := make(map[int]*element)
	for i, num := range nums {
		if _, ok := hash[num]; !ok {
			hash[num] = &element{
				start: i,
				end:   i,
				count: 1,
			}
			continue
		}

		hash[num].end = i
		hash[num].count++
		// 计算当前数组的度
		if hash[num].count > maxCount {
			maxCount = hash[num].count
		}
	}

	ans := len(nums)
	// 遍历无序集合，找到符合最大的度的元素，并计算其最短连续子数组长度
	for _, ele := range hash {
		if ele.count == maxCount {
			ans = min(ans, ele.end-ele.start+1)
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
