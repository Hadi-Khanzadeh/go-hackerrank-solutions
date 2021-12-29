package main

import (
	"fmt"
)

func main() {

	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for j := 0; j < m; j++ {
		fmt.Scan(&b[j])
	}

	var c int

	for i := 1; i <= 100; i++ {
		factor := true
		for j := 0; j < n; j++ {
			if i%a[j] != 0 {
				factor = false
				break
			}
		}
		if factor {
			for j := 0; j < m; j++ {
				if b[j]%i != 0 {
					factor = false
					break
				}
			}
		}
		if factor {
			c++
		}
	}
	fmt.Println(c)
}
