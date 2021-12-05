package main

// 时间复杂度 O(n)，n 为字符串 s 长度。因为需要遍历字符串 s，字符串 s 长度为 n。
// 空间复杂度 O(n)，因为 go 中字符串是不可变的，所以需要转为字节数组才能修改，字节数组长度取决于字符串 s 的长度。
func toLowerCase(s string) string {
	ans := []byte(s)
	for i, v := range ans {
		if 'A' <= v && v <= 'Z' {
			ans[i] = v + 'a' - 'A'
		}
	}
	return string(ans)
}
