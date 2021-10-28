package main

// 时间复杂度为 O(n+m), n 为课程总数量， m 为有依赖的数组的长度（prerequisites 的长度），即借助了邻接表的广度优先遍历的时间复杂度。
// 空间复杂度为 O(n+m), n 为课程总数量， m 为有依赖的数组的长度（prerequisites 的长度），因为借助了一个邻接表（n+m），一个入度数组(n)，一个队列(n)，所以总空间复杂度为 O(n+m)。
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		// 统计元素入度
		in = make([]int, numCourses)
		// 统计元素的依赖顺序，只有下表对应的课程入度为 0 才能执行 val
		// 即邻接表
		to = make([][]int, numCourses)
	)
	// 遍历课程信息
	for _, v := range prerequisites {
		// a 依赖于 b
		a, b := v[0], v[1]
		// 计算 a 的入度
		in[a]++
		to[b] = append(to[b], a)
	}

	// 定义一个队列
	q := make([]int, 0, numCourses)
	// 将所有入度为 0 的课程放入队列
	for i, v := range in {
		if v == 0 {
			q = append(q, i)
		}
	}

	// 能被执行的课程
	lessons := make([]int, 0, numCourses)
	// 遍历队列
	for len(q) > 0 {
		// 弹出队头
		front := q[0]
		// 队头出队
		q = q[1:]
		// 更新能被执行的的课程
		lessons = append(lessons, front)

		for _, v := range to[front] {
			// 入度减一，因为所依赖的某个课程已经执行了
			in[v]--
			// 如果入度边为 0，表示没有依赖了，就放入队列执行
			if in[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(lessons) != numCourses {
		return nil
	}
	return lessons
}
