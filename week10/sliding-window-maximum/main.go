package main

import (
	"container/heap"
	"sort"
)

// 使用一个全局变量 a 保存 nums，主要用作实现堆的比较函数 Less。
var a []int

// hp 定义大根堆的实现结构，堆中保存数组下标。
type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	tmp := h.IntSlice
	v := tmp[len(a)-1]
	h.IntSlice = tmp[:len(tmp)-1]
	return v
}

// 时间复杂度：O(nlogk)，n 为 nums 长度。因为需要遍历所以元素才能得到答案，每遍历一个元素几乎都要执行一个入堆和出堆，这些操作的复杂度为 logk（因为 堆容量为 k），所以时间复杂度为 O(nlogk)。
// 空间复杂度：O(n)，因为用了一个长度为 n 的全局全局变量保存了 nums，一个长度为 k 的堆辅助计算，所以空间复杂度为 O(n + k)，也就是 O(n)。
func maxSlidingWindow(nums []int, k int) []int {
	// 将 nums 保存到全局变量
	a = nums
	// 定一个长度为 k 的堆
	q := &hp{make([]int, k)}
	// 将数组中前 k 个元素的下标保存到堆
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	// 初始化堆
	heap.Init(q)

	n := len(nums)
	// 定义答案长度，因为窗口长度为 k，所以除了最开始的 k - 1 个元素不会输出窗口最大值，之后的元素都会输出，所以答案长度为 n-(k-1)
	ans := make([]int, 1, n-k+1)
	// 因为前 k 个元素已经入堆，所以此时堆顶元素就是最大值，而堆顶就是数组的第一个元素
	ans[0] = nums[q.IntSlice[0]]
	// 将 k 之后的元素入堆
	for i := k; i < n; i++ {
		heap.Push(q, i)
		// 如果堆顶已经不在窗口内了，弹出堆顶
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		// 此时的堆顶就是窗口最大值
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
