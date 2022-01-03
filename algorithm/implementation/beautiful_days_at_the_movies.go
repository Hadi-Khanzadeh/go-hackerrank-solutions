package main

import (
	"fmt"
)

func main() {
	var i, j, k, n int
	fmt.Scan(&i, &j, &k)
	for x := i; x <= j; x++ {
		r := Reverse(x)
		v := x - r
		if v%k == 0 {
			n++
		}
	}
	fmt.Println(n)
}

func Reverse(x int) int {
	result := 0
	for x/10 != 0 {
		if x%10 != 0 {
			result += (x % 10)
		}
		result *= 10
		x /= 10
	}
	result += x
	return result
}
