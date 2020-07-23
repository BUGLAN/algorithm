package algorithm

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
	res := make([]string, 0)
	fullArrange([]rune(s), 0, &res)
	return res
}

func fullArrange(runes []rune, start int, res *[]string) {
	if runes == nil {
		return
	}

	if start == len(runes)-1 {
		*res = append(*res, string(runes))
	} else {
		for i := start; i < len(runes); i++ {
			// 交换
			if i != start {
				runes[i], runes[start] = runes[start], runes[i]
			}
			fullArrange(runes, start+1, res)
			if i != start {
				runes[start], runes[i] = runes[i], runes[start]
			}

		}
	}
}
