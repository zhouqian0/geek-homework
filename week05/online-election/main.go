package main

type TopVotedCandidate struct {
	winners []int // 保存根据构建时传入的 persons 统计出的每个时刻的 winner
	times   []int // 保存构建对象时的 time 数组
}

// 时间复杂度 O(n), n 为 persons 的长度。
// 空间复杂度 O(n)，因为使用了一个无序集合保存每个候选人的票数，而候选人数量不超过 n，除此之外只使用了一个常数级别的变量，所以空间复杂度为 O(n)。
func Constructor(persons []int, times []int) TopVotedCandidate {
	t := TopVotedCandidate{
		winners: make([]int, len(persons)),
		times:   times,
	}

	// 定义一个无序集合保存每个参选人的票数
	mp := make(map[int]int)
	// 当前的最高票数
	maxVote := -1
	for i := range persons {
		mp[persons[i]]++
		// 如果当前参选人的票数大于此前的最高票数
		if mp[persons[i]] >= maxVote {
			// 更新最高票数
			maxVote = mp[persons[i]]
			// 当下的 winner 就是就是当前参与人
			t.winners[i] = persons[i]
		} else {
			// 如果当前参选人的票数低于此前的最高票数，那么当下的 winner 等于 for 循环中的上一个元素的 winner
			t.winners[i] = t.winners[i-1]
		}
	}
	return t
}

// 时间复杂度为 O(logn)，n 为 times 的长度，主要使用了二分的思想在 times 里找出最后一个小于等于 t 的元素，而二分的时间复杂度为 O(logn)，所以总时间复杂度为 O(logn)。
// 空间复杂度 O(1)，因为只借助了常数级别的辅助变量辅助二分计算，所以空间复杂度为 O(1)。
func (this *TopVotedCandidate) Q(t int) int {
	// 使用二分思想，在 times 里找出最后一个小于等于 t 的元素
	left, right := 0, len(this.times)-1
	for left < right {
		mid := (left + right + 1) / 2
		if this.times[mid] <= t {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// 因为 winner 里已经统计好了每个时刻的 winner
	return this.winners[right]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
