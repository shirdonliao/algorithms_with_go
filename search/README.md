# 二分查找算法

二分查找容易理解。对有序集合`[left, right]` 取其中点 `[left, ..., mid, ..., right]`，分三种情况：

- 如果要查找的值 `v < mid`，则说明 v 在 `[left, ..., mid)` 区间。继续二分查找
- 如果要查找的值 `v > mid`，则说明 v 在 `(mid, ..., right]` 区间。继续二分查找
- 相等则已找到

## 查找过程

对数组 `[1, 3, 4, 6, 7, 8, 10, 13, 14]` 查找元素 4，其 mid 取值如下

```shell
binary_search git:(master) ✗ go run main.go
[DEBUG arr[mid]]:        7	# 4 < 7 往左边走
[DEBUG arr[mid]]:        3	# 4 > 3 往右边走
[DEBUG arr[mid]]:        4	# v 与 arr[mid] 相等，已找到
```

## 查找效果

 ![](http://p7f8yck57.bkt.clouddn.com/2018-06-15-114725.jpg)

## 复杂度

### 时间复杂度

二分查找每次会减半搜索区域，时间复杂度为 **O(log N)**

### 空间复杂度

无需借助外部空间，空间复杂度为 **O(1)**



## 使用场景

有序集合的元素查找。

需注意，取 `mid` 一般会直接使用 `mid = (left + right) / 2`，当处理大数组时 left + right 可能会整型上溢，应该使用 `mid := left + ((right - left) >> 1)` 来取中点值。
