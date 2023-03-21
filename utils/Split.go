package utils

// Split 字符串截取。中文字符算一个长度，如果总长度超过 length，那么截取至 length 长度字符串
func Split(s string, length int) string {
	if len([]rune(s)) > length {
		s = string([]rune(s)[0:length])
	}
	return s
}
