package main

// 时间复杂度 O(n)，n 为字符串 s 长度。最多需要遍历整个字符串才能获得结果。
// 空间复杂度 O(1)，因为只借助了常数级别的辅助变量。
func lengthOfLastWord(s string) (ans int) {
	idx := len(s) - 1
	for idx >= 0 && s[idx] == ' ' {
		idx--
	}
	for idx >= 0 && s[idx] != ' ' {
		idx--
		ans++
	}
	return
}
