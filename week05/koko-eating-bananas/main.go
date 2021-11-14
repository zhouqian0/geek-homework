package mai

import "math"

// 时间复杂度为 O(nlogm)，其中 n 为 piles 长度， m 为 piles 中的所有堆的香蕉总量减去 piles中最小一堆中包含的香蕉总量。
// 		因为在刚开始执行二分时，所有的待二分的元素总量为 m，所以二分耗时 O(logm)，
//		而每次的二分都要执行 valid 函数，在 valid 函数中，需要遍历一次 piles，而遍历 piles 的时间复杂度为 O(n)，
//		所以总时间复杂度为 O(nlogm)。
// 空间复杂度为 O(1)，因为只借助了常数级别的辅助变量用作辅助二分计算，所以总空间复杂度为 O(1)。
func minEatingSpeed(piles []int, h int) int {
	// 定义二分边界
	// 左边界为 piles 中最小堆的数量，右边界为 piles 中所有香蕉总量
	var left, right int
	for _, pile := range piles {
		left = min(left, pile)
		right += pile
	}

	// 开始二分，找到第一个满足条件的 mid 值（mid 表示每小时吃 mid 根香蕉）
	for left < right {
		mid := (left + right) / 2
		// 如果当前的速度不满足条件，则说明需要加快速度（mid 表示每小时吃 mid 根香蕉）
		if !valid(piles, h, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

// valid 能否每小时吃 num 根香蕉，持续 h 小时的场景下，吃完 piles 的香蕉。
func valid(piles []int, h, num int) bool {
	// 如果传入每小时吃 0 根香蕉，就返回 false
	if num == 0 {
		return false
	}

	// 计算按照 num 根香蕉/小时 的速度没需要多久吃完 piles 中的香蕉
	hour := 0
	for _, pile := range piles {
		// 向上取整，因为当堆中的香蕉数量少于 num 时，也会单独耗去一个小时
		hour += int(math.Ceil(float64(pile) / float64(num)))
	}
	// 返回能否在规定的 h 小时内吃完香蕉
	return hour <= h
}

// min 比较两个整数，返回其中较小值。
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
