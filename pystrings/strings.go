package pystrings

import (
	"strings"
)

func RSplit(s string, sub string) []string {
	return RSplitN(s, sub, -1)
}

func RSplitN(s string, sub string, n int) []string {
	var subs []string
	var end bool
	for i := 0; i < n; i++ {
		index := strings.LastIndex(s, sub)
		if index < 0 {
			subs = append(subs, s)
			end = true
			break
		} else {
			subs = append(subs, s[index+len(sub):])
			s = s[:index]
		}
	}
	if !end { //还有未完成的直接加入
		subs = append(subs, s)
	}

	// reverse
	for i := 0; i < len(subs)/2; i++ {
		last := len(subs) - 1 - i
		subs[i], subs[last] = subs[last], subs[i]
	}
	return subs
}

// 实现字符串数组的高阶函数
func Filter(fn func(x string) bool, in []string) []string {
	out := make([]string, len(in))
	var index int
	for i := 0; i < len(in); i++ {
		if fn(in[i]) {
			out[index] = in[i]
			index++
		}
	}
	return out[:index]
}

func Generator(fn func(x string) string, in []string) []string {
	out := make([]string, len(in))
	var index int
	for i := 0; i < len(in); i++ {
		out[index] = fn(in[i])
		index++
	}
	return out[:index]
}
