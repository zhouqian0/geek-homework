package main

type pos struct {
	count int
	start int
	end   int
}

func findShortestSubArray(nums []int) (ans int) {
	hash := make(map[int]*pos)
	for i, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num].count++
			hash[num].end = i
			continue
		}
		hash[num] = &pos{
			count: 1,
			start: i,
			end:   i,
		}
	}

	max := 0
	ans = 49999
	for _, num := range nums {
		if max < hash[num].count {
			max = hash[num].count
			ans = hash[num].end - hash[num].start + 1
		} else if max == hash[num].count {
			ans = min(ans, hash[num].end-hash[num].start+1)
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
