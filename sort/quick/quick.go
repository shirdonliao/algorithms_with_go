package main

import (
	"../../common"
	"fmt"
)

func main() {

	arr := common.GetArr(5, 20)
	//arr = []int{27, 38, 12, 39, 27, 16}
	fmt.Println("[UNSORTED]:\t", arr)
	fmt.Println("[SORTED]:\t", quickSort(arr))
}

func quickSort(arr []int) []int {

	n := len(arr)
	// 递归结束条件
	if n <= 1 {
		temp := make([]int, n)
		copy(temp, arr)
		return temp
	}

	// 使用第一个元素作为基准值
	pivot := arr[0]

	// 小元素 和 大元素各成一个数组
	low := make([]int, 0, n)
	high := make([]int, 0, n)

	// 更小的元素放 low[]
	// 更大的元素放 high[]
	for i := 1; i < n; i++ {
		if arr[i] < pivot {
			low = append(low, arr[i])
		} else {
			high = append(high, arr[i])
		}
	}
	// 子区间递归快排，分治排序
	low, high = quickSort(low), quickSort(high)
	fmt.Println("[DEBUG low]:\t", low)
	fmt.Println("[DEBUG high]:\t", high)
	return append(append(low, arr[0]), high...)
}
