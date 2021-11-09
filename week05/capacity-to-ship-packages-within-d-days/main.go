package main

// 时间复杂度 O(logn)，主要还是二分思想，二分的时间复杂度为 O(logn)，但如果所有元素都重复的话，时间复杂度会降低到 O(n)，因为需要一直执行 r--，直到 r == 0。
// 空间复杂度 O(1)，因为只使用了常数级别的辅助变量。
func findMin(nums []int) int {
	// 定义二分的左右边界
	l, r := 0, len(nums)-1
	for l < r {
		// 取二分中间点
		mid := (l + r) / 2
		// 将中间点和右边界点，做判断，总共有三种场景
		// 1. mid 点大于右边界，如 [3, 3, 1], nums[1] > nums[2]。此时，将左边界移动到中间点以后。
		// 2. mid 点小于右边界，如 [3, 1, 3]，nums[1] < nums[2]。此时，将右边界等于中间点。
		// 3. mid 点等于右边界，如 [1, 3, 3]，nums[1] == nums[2]。此时将右边界左移一位，避开重复的元素。
		switch {
		case nums[mid] > nums[r]:
			l = mid + 1
		case nums[mid] < nums[r]: // mid 可能是最小值
			r = mid
		default: // 因为 mid 和 r 值相同，mid 可能是最小值，并且不确定 mid 和 r 执行是否出现断层，如 [7, 7, 7, 2, 7]，所以移动右边界
			r--
		}
	}
	// 此时返回左边界和右边界都可以，因为 l == r（这是跳出上面 for 循环的条件）
	return nums[r]
}
