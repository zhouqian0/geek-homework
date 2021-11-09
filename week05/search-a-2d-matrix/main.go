package main

// 时间复杂度：O(logm*n)，m, n 分别为矩阵行数和列数，借助二分需要 logm 的时间去找到目标行，再用 logn 的时间在目标行中二分找到目标列
// 		所以时间复杂度为 O(logn + logm)，也就是 O(logm*n)。
// 空间复杂度：O(1)，因为只借助了常数级别的辅助变量用作二分。
func searchMatrix(matrix [][]int, target int) bool {
	// 获取行数列数
	m, n := len(matrix), len(matrix[0])

	// 通过二分确定行数，因为根据题意 每行中的整数从左到右按升序排列，每行的第一个整数大于前一行的最后一个整数
	l, r := 0, m-1
	for l <= r {
		rowMid := (l + r) / 2
		switch {
		case matrix[rowMid][0] > target: // rowMid 行的第一个元素大于 target，说明 target 在之前的行
			r = rowMid - 1
		case matrix[rowMid][n-1] < target: // rowMid 行的最后一个元素的值小于 targett，说明 target 在之后的行
			l = rowMid + 1
		default: // 找到了 target 所在行
			// 开始对列二分
			l, r := 0, n-1
			for l <= r {
				colMid := (l + r) / 2
				switch {
				case matrix[rowMid][colMid] > target: // rowMid 行的 colMid 元素大于 target，说明 target 在之前的列
					r = colMid - 1
				case matrix[rowMid][colMid] < target: // rowMid 行的 colMid 元素小于 target，说明 target 在之后的列
					l = colMid + 1
				default:
					return true
				}
			}
			// 没有找到 target 所在列
			return false
		}
	}
	// 没有找到 target 所在行
	return false
}
