package main

// 时间复杂度 O(nlogm)，n 为 weights 长度，m 为 weights 中所有元素的和减去最大的那个元素的值的结果。
// 		其中 m 是执行二分时元素的数量，所以二分操作的时间消耗为 O(logm)
//		因为每执行一次二分都需要执行一次 valid 函数，而 valid 函数需要遍历一次 weights，所以时间消耗 O(n)
//		综合在一起就是，总时间复杂度就是 O(nlogm)。
// 空间复杂度 O(1)，主要借助二分答案思想，使用了常数级别的辅助变量，所以空间复杂度为 O(1)。
func shipWithinDays(weights []int, days int) int {
	// 定义二分左右边界，左边界为最重的一个包裹的重量，右边界为所有包裹的合计重量
	var left, right int
	for _, weight := range weights {
		left = max(left, weight)
		right += weight
	}
	// 开始二分
	for left < right {
		mid := (left + right) / 2
		// 判断如果使用 mid 作为载重，能否满足要求
		if !valid(weights, days, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func valid(weights []int, days int, carry int) bool {
	// 按照给定负载运所有货物需要的天数
	day := 1
	// 某一天已运载货物重量，用作临时变量，方便判断当前的货物是否需要隔天运送
	cnt := 0
	for _, weight := range weights {
		// 如果当前 day 可以运载下货物 weight
		if cnt+weight <= carry {
			// 运载货物
			cnt += weight
		} else { // 当前 day 不可以运载下货物 weight
			// 换下一天运载货物 weight
			day++
			// 重置 day 天已运载的货物重量
			cnt = weight
		}
	}
	// 判断按照负载 carry 运载货物能否在 days 天内完成
	return day <= days
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
