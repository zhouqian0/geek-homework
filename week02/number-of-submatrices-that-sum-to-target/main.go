// 力扣 1074 元素和为目标值的子矩阵数量（https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/）
package main

// 时间复杂度 O(n*n*m) n 为矩阵行数，m 为矩阵列数。 因为总共有三个 for 循环，且都是嵌套执行，前两个耗时 n，最后一个耗时 m，故而总时间复杂度 O(n*n*m)。
// 空间复杂度 o(m)，m 为矩阵列数。因为申请了一个数组来保存每列的和，虽然这个数组出现在第一个 for 循环中，但每次第一个 for 循环执行时，只是重置了它的值而已。
func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	// 枚举上边界
	for i := range matrix {
		sum := make([]int, len(matrix[0]))
		// 枚举下边界
		for _, row := range matrix[i:] {
			// 计算每列的和，这里将上下边界内的矩阵压缩成了一维数组
			for c, v := range row {
				sum[c] += v
			}
			ans += subarraySum(sum, target)
		}
	}
	return
}

// subarraySum 从一个一维数组中查找和为 k 的子数组。
func subarraySum(nums []int, k int) (ans int) {
	// 用一个无序集合保存前缀和及其出现次数
	hash := map[int]int{0: 1}
	//前缀和
	pre := 0
	for _, v := range nums {
		// 计算前缀和
		pre += v
		// 判断是否有满足条件的子数组
		if v, ok := hash[pre-k]; ok {
			// 更新计数
			ans += v
		}
		// 更新当前缀和的出现次数
		hash[pre]++
	}
	return
}
