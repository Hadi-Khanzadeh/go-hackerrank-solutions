package main

import (
	"fmt"
)

func main() {
	var v, n int
	fmt.Scan(&v, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] == v {
			fmt.Println(i)
		}
	}
}
