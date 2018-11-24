package main

import (
	"fmt"
	//"../common"
)

func main() {
	//arr := common.GetArr(3, 3)
	arr := []int{1, 2, 3}
	fmt.Println("[UNSORTED]:  ", arr)

	n := len(arr)
	// 遍历所有元素
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// 左元素 > 右元素
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println("[DEBUG]:  ", arr)
		}
	}

	fmt.Println("[SORTED]:  ", arr)
}
