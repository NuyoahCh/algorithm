package main

import "fmt"

func main() {
	// 创建动态数组
	// 不用显式指定数组大小，它会根据实际存储的元素数量自动扩缩容
	arr := make([]int, 0)

	for i := 0; i < 10; i++ {
		// 在末尾追加元素，时间复杂度 O(1)
		arr = append(arr, i)
	}

	// 在中间插入元素，时间复杂度 O(N)
	// 在索引 2 的位置插入元素 666
	arr = append(arr[:2], append([]int{666}, arr[2:]...)...)

	// 在头部插入元素，时间复杂度 O(N)
	arr = append([]int{-1}, arr...)

	// 删除末尾元素，时间复杂度 O(1)
	arr = arr[:len(arr)-1]

	// 删除中间元素，时间复杂度 O(N)
	// 删除索引 2 的元素
	arr = append(arr[:2], arr[3:]...)

	// 根据索引查询元素，时间复杂度 O(1)
	a := arr[0]
	fmt.Println(a)

	// 根据索引修改元素，时间复杂度 O(1)
	arr[0] = 100

	// 根据元素值查找索引，时间复杂度 O(N)
	index := -1
	fmt.Println(index)

	for i, v := range arr {
		if v == 666 {
			index = i
			break
		}
	}
}
