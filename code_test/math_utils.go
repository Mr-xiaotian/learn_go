package main

import "fmt"

func sumEven(num int) int {
	sum_even_num := 0
	for index := range num + 1 {
		if index%2 != 0 {
			continue
		}
		sum_even_num += index
	}
	return sum_even_num
}

func countOdd(nums []int) int {
	count_odd_num := 0
	for _, num := range nums {
		if num%2 == 0 {
			continue
		}
		count_odd_num += 1
	}
	return count_odd_num
}

func printTable(n int) {
	for index_0 := range n {
		x := index_0 + 1
		for index_1 := range x {
			y := index_1 + 1
			fmt.Printf("%dx%d=%d ", y, x, y*x)
		}
		fmt.Println()
	}
}
