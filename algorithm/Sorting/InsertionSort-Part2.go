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

func insertionSort2(n int, arr []int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]

			}
			j = j - 1
		}
		PrintArr(arr, n)
	}

}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	insertionSort2(n, arr)

}
