package main

import (
	"fmt"
)

func bubbleSort(arr []int, l int, r int) []int {

	subarray := arr[l : r+1]

	// 使用冒泡排序对子数组进行排序
	for i := 1; i < len(subarray); i++ {
		for j := 0; j < len(subarray)-i; j++ {
			if subarray[j] > subarray[j+1] {
				subarray[j], subarray[j+1] = subarray[j+1], subarray[j]
			}
		}
	}
	// 将排序后的子数组替换回原数组
	copy(arr[l:r+1], subarray)
	return arr
}
func Sort(array []int, l int, r int) {

	array = bubbleSort(array, l, r)
	for _, v := range array {
		fmt.Print(v, " ")
	}

}
func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)
	var array = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])

	}
	Sort(array, l, r)
}
