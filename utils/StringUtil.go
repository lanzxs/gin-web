package utils

import "strings"

// 是否空字符串
func StrIsEmpty(str string) bool {
	return str == "null" || strings.TrimSpace(str) == ""
}
