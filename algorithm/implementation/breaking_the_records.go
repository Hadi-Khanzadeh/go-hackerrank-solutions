package main

import (
	"fmt"
)

func main() {
	var n, min, max, nmin, nmax int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		if i == 0 {
			min = t
			max = t
		} else {
			if t < min {
				nmin++
				min = t
			}
			if t > max {
				nmax++
				max = t
			}
		}
	}
	fmt.Println(nmax, nmin)
}
