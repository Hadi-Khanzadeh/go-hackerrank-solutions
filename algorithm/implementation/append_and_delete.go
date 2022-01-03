package main

import (
	"fmt"
)

func main() {
	var s, t string
	var k int
	fmt.Scan(&s)
	fmt.Scan(&t)
	fmt.Scan(&k)

	index := 0
	for i := 0; i < len(s) && i < len(t); i++ {
		if s[i] != t[i] {
			break
		} else {
			index++
		}
	}
	sum_len := len(s) + len(t)
	if sum_len-2*index > k {
		fmt.Println("No")
	} else if (sum_len-2*index)%2 == k%2 {
		fmt.Println("Yes")
	} else if sum_len-k < 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
