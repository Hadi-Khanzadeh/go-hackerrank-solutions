package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	socks := make(map[int]int)

	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)

		socks[t]++
	}

	count := 0

	for _, v := range socks {
		count += v / 2
	}

	fmt.Println(count)
}
