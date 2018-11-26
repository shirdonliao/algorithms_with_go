#快速排序法

快速排序是冒泡排序的优化版，同属于交换类排序。它使用了 “分治法” 的思想，将集合分割为相似子集，再对子集进行递归排序，最后将子集的排序结果组合即可。

##排序过程
quick_sort $ go run main.go
[UNSORTED]:      [27 38 12 39 27 16]
[DEBUG low]:     []
[DEBUG high]:    [16]
[DEBUG low]:     [27]
[DEBUG high]:    [39]
[DEBUG low]:     [12 16]
[DEBUG high]:    [27 38 39]
[SORTED]:        [12 16 27 27 38 39]

##复杂度
###时间复杂度
最好情况：递归求证，O( NlogN )

最坏情况：退化为冒泡排序 O( N^2 )

平均复杂度：O(NlogN)

##使用场景
适用于大规模无序数据排序。
##总结
快排的关键是“分治”的思想，对应到代码实现即递归。需注意递归退出 return 的条件，本例是子集中只有 1 个元素时视为排序完毕，递归结束。快排还有其他两种优化实现，可参考讨论：how-to-optimize-quicksort