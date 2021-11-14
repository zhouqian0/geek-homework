package main

// 时间复杂度 O(nlogn)，n 为 nums 长度，主要采用了归并排序的思想，划分子数组的次数为 logn 次，每次需要遍历所有元素，来使其有序，耗时 n，所以总时间复杂度为 O(nlogn)。
// 空间复杂度 O(n)，n 为 nums 长度，前缀和数组需要 O(n) 的空间，归并排序的递归栈深度为 logn，每次递归都需要一些额外的空间保存排序后的子数组，子数组元素总数量为 n，所以总空间复杂度为 O(n)。（logn < n）
func countRangeSum(nums []int, lower int, upper int) int {
	// 借助归并排序思想
	var mergeCount func(arr []int) int
	mergeCount = func(arr []int) (cnt int) {
		// 因为数组中元素少于两个，就构成不了区间
		if len(arr) < 2 {
			return 0
		}

		// 将 nums 拆分成两个子数组
		n1 := append([]int{}, arr[:len(arr)/2]...)
		n2 := append([]int{}, arr[len(arr)/2:]...)
		// 对子数组归并排序，并统计子数组中符合条件的区间和个数
		cnt = mergeCount(n1) + mergeCount(n2)

		// 在排序后的 n1 和 n2 数组中，找到符合条件的区间和个数
		// 每遍历一个 n1 中的元素，就在 n2 中寻找是否有满足条件的区间和
		l, r := 0, 0
		for _, num := range n1 {
			// 因为 n2 内元素已经排序好了，所以直接进行区间和计算
			// 找到第一个大于等于 low 的元素
			for l < len(n2) && n2[l]-num < lower {
				l++
			}
			// 找到第一个大于 upper 的元素
			for r < len(n2) && n2[r]-num <= upper {
				r++
			}
			// 因为右边界 r 是一个大于 upper 的第一个元素，r - l 就是满足条件区间和个数
			cnt += (r - l)
		}

		// 排序 arr，使其有序
		// 此时 n1 和 n2 都是有序数组
		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(n1) && (p2 == len(n2) || n1[p1] < n2[p2]) {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}
		// 返回统计出的区间和个数
		return
	}

	// 计算前缀和（第一个元素为 0）
	perSum := make([]int, len(nums)+1)
	for i := range nums {
		perSum[i+1] = perSum[i] + nums[i]
	}
	// 对前缀和数组进行区间和个数统计
	return mergeCount(perSum)
}
