package main

import "fmt"

func sumEven(num int) int {
	sum_even_num := 0
	for index := range num + 1 {
		// fmt.Println(index)
		if index%2 != 0 {
			continue
		}
		sum_even_num += index
	}
	return sum_even_num
}

func reverseArray(arr []int) []int {
	var reverse_array []int
	arr_max_index := len(arr) - 1

	for index := range arr {
		reverse_array = append(reverse_array, arr[arr_max_index-index])
	}
	return reverse_array
}

func countOdd(nums []int) int {
	count_odd_num := 0
	for _, num := range nums {
		if num%2 == 0 {
			continue
		}
		// fmt.Println(num)
		count_odd_num += 1
	}
	return count_odd_num
}

func reverseString(str string) string {
	rune_str := []rune(str)
	reverse_string_rune := []rune{}
	max_str_index := len(rune_str) - 1

	for index := range rune_str {
		reverse_string_rune = append(reverse_string_rune, rune_str[max_str_index-index])
	}
	return string(reverse_string_rune)
}

func isPalindrome(s string) bool {
	return reverseString(s) == s
}

func printTable(n int) {

}

func main() {
	fmt.Println(sumEven(6))
	fmt.Println(reverseArray([]int{1, 2, 3, 4}))
	fmt.Println(countOdd([]int{1, 2, 3, 4, 5}))
	fmt.Println(isPalindrome("level"))
}
