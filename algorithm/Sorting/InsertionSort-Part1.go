package main

import (
	"fmt"
)

func PrintArr(arr1 []int, n int) {
	var i int = 0
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", arr1[i])
	}
	fmt.Println()
}

func insertionSort1(n int, arr []int) {
	l, i := arr[n-1], 0
	for i = int(n) - 2; arr[i] >= l; i-- {
		arr[i+1] = arr[i]
		PrintArr(arr, n)
		if i == 0 {
			i--
			break
		}
	}
	arr[i+1] = l
	PrintArr(arr, n)
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	insertionSort1(n, arr)

}
