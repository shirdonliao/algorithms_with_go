package main

import (
	"github.com/shirdonliao/algorithms_with_go/common"
	"fmt"
)

func main() {
	arr := common.GetArr(5, 20)
	arr = []int{29, 10, 14, 37, 14}
	fmt.Println("[UNSORTED_ARR]: ", arr)

	n := len(arr)
	if n <= 1 {
		fmt.Println("[ALREADY SORTED]: ", arr)
		return
	}

	// 遍历所有元素
	for i := 1; i < n; i++ {
		// 向前找位置
		for j := i; j > 0; j-- {
			// 合适位置插入
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		fmt.Println("[DEBUGING]: ", arr)
	}
	fmt.Println("[SORTED_ARR]: ", arr)
}
