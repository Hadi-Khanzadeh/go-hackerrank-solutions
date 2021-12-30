package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	c := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if (a[i]+a[j])%k == 0 && i < j {
				c++
			}
		}
	}
	fmt.Println(c)
}
