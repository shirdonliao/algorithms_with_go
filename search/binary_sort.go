package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	fmt.Println(binarySearch(arr, 4)) // 4
}

//
// 在已排序数组 sortedArr 中搜索 v
// 存在则返回索引，否则返回 -1
//
func binarySearch(sortedArr []int, v int) int {
	// 左右两个游标
	left, right := 0, len(sortedArr)-1

	// 游标未相遇
	for left < right {
		// mid := (right - left) / 2
		// 避免整型上溢
		mid := left + ((right - left) >> 1)
		fmt.Println("[DEBUG arr[mid]]:\t", sortedArr[mid])
		if sortedArr[mid] > v {
			// 中点值更大，则在左区间
			right = mid - 1
		} else if sortedArr[mid] < v {
			// 中点值更小，则在右区间
			left = mid + 1
		} else {
			// 相等则已找到
			return mid
		}
	}
	// 未找到
	return -1
}
