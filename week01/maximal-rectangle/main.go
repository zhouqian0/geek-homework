package main

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return 0
	}

	// 计划一层一层的计算最大矩形面积
	heights := make([]int, len(matrix[0]))
	// 遍历整个矩阵
	for _, rows := range matrix {
		for i, v := range rows {
			// 题目规则 matrix[i][j] 为 '0' 或 '1'
			if v == '1' {
				// 字符 '1' 在上下两层中同一下标处连续
				heights[i]++
			} else {
				// 遇到 '1' 不连续的情况重新开始计算
				heights[i] = 0
			}
		}

		// 在每层统计完成后计算当前最优解
		ans = max(ans, largestRectangleArea(heights))
	}

	return ans
}

type rect struct {
	height int
	width  int
}

func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)

	// 维护一个单调栈，递增
	stack := make([]rect, 0, n)
	// 在数组末尾增加一个元素，保证在遍历到最后一个元素时，破坏单调递增，使所有元素出栈计算
	heights = append(heights, 0)

	// 遍历所有元素
	for _, height := range heights {
		accumulatedWidth := 0

		// 如果栈中已有元素，检查栈顶元素和当前元素是否满足单调增
		for len(stack) > 0 {
			// 获取栈顶元素
			top := stack[len(stack)-1]
			// 满足单调增则跳出
			if top.height <= height {
				break
			}

			// 因当前元素破坏了单调增，弹出栈顶元素，累计宽度并更新答案
			stack = stack[:len(stack)-1]
			accumulatedWidth += top.width
			ans = max(ans, accumulatedWidth*top.height)
		}

		// 新元素入栈
		stack = append(stack, rect{
			height: height,
			width:  accumulatedWidth + 1,
		})
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
