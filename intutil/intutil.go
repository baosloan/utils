package intutil

import "math"

// Intersect 求两个整形切片的交集
func Intersect(slice1, slice2 []int) []int {
	m := make(map[int]struct{})
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	var result []int
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

// Reverse 将整形切片翻转(不改变原有切片)
func Reverse(slice []int) []int {
	var length = len(slice)
	var result = make([]int, length)
	middle := int(math.Round(float64(length) / 2))
	for i := 0; i < middle; i++ {
		result[i], result[length-i-1] = slice[length-i-1], slice[i]
	}
	return result
}
