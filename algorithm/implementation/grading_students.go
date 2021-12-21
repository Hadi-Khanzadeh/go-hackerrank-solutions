package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)
	numbers := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&numbers[i])
		if numbers[i] < 38 || (numbers[i]%5) < 3 {
			fmt.Println(numbers[i])
		} else {
			fmt.Println((numbers[i]/5 + 1) * 5)
		}
	}
}
