package stringutil

import "math"

// Intersect 求两个字符串切片的交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]struct{})
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	var result []string
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

// Reverse 将字符串切片翻转(不改变原有切片)
func Reverse(slice []string) []string {
	var length = len(slice)
	var result = make([]string, length)
	middle := int(math.Round(float64(length) / 2))
	for i := 0; i < middle; i++ {
		result[i], result[length-i-1] = slice[length-i-1], slice[i]
	}
	return result
}
