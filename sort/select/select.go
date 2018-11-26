package main

import (
	"../../common"
	"fmt"
)

func main() {
	arr := common.GetArr(5, 20)
	//arr = []int{29, 10, 14, 37, 13}
	fmt.Println("[UNSORTED]:\t", arr)

	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 假设无序区间第一个值为最小值
		min := i
		for next := min + 1; next < n; next++ {
			// 找到更小值，记录其位置
			if arr[min] > arr[next] {
				min = next
			}
		}
		// 将无序区间的最小值追加到有序区间
		fmt.Println("[DEBUG min]:\t", arr[min])
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println("[SORTED]:\t", arr)
}
