package main

import (
	"fmt"
)

func main() {

	var s, t, a, b, m, n int
	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&m, &n)
	var anum, bnum, d int
	for i := 0; i < m; i++ {
		fmt.Scan(&d)
		if a+d >= s && a+d <= t {
			anum++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		if b+d >= s && b+d <= t {
			bnum++
		}
	}
	fmt.Printf("%d\n%d", anum, bnum)
}
