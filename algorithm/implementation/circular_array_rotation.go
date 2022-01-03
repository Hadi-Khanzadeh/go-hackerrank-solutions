package main

import (
	"fmt"
)

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var query int
	for i := 0; i < q; i++ {
		fmt.Scan(&query)
		fmt.Println(a[(n-k+query)%n])
	}
}
