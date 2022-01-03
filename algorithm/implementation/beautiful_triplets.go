package main

import (
	"fmt"
)

func main() {
	var n, d, result int
	fmt.Scan(&n, &d)
	a := make([]int, n)
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		m[a[i]] = i
	}
	for i := 0; i < n-2; i++ {
		aj := d + a[i]
		if j, found := m[aj]; found && i < j {
			ak := d + a[j]
			if k, found := m[ak]; found && j < k {
				result++
			}
		}
	}
	fmt.Println(result)
}
