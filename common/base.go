package common

import (
	"math/rand"
	"time"
)

func GetArr(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max + 1)
	}
	return arr
}

//
// 获取 [min, max] 的连续数值数组
//
func GetRange(min, max int) []int {
	arr := make([]int, max-min+1)
	for i := range arr {
		arr[i] = min + i
	}
	return arr
}
