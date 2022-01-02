package main

import (
	"fmt"
)

func main() {

	var t int
	fmt.Scan(&t)
	for j := 0; j < t; j++ {
		var n, k, x, count int
		fmt.Scan(&n, &k)
		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			if x <= 0 {
				count++
			}
		}
		if count < k {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
