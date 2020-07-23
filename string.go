package algorithm

import (
	"fmt"
)

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)

	i, j := 0, len(s)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

// FullArrange 全排列
func FullArrange(s string) []string {
	fullArrange([]rune(s), 0)
	return []string{}
}

func fullArrange(runes []rune, start int) {
	if runes == nil {
		return
	}

	if start == len(runes)-1 {
		fmt.Println(string(runes))
	} else {
		for i := start; i < len(runes); i++ {
			runes[i], runes[start] = runes[start], runes[i]
			fullArrange(runes, start+1)
			runes[start], runes[i] = runes[i], runes[start]
		}
	}
}
