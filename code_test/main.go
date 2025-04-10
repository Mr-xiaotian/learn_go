package main

import (
	"fmt"
	"math"
)

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
	for index_0 := range n {
		x := index_0 + 1
		for index_1 := range x {
			y := index_1 + 1
			fmt.Printf("%dx%d=%d ", y, x, y*x)
		}
		fmt.Println()
	}
}

func findMaxKey(m map[string]int) string {
	var max_key string
	max_value := math.MinInt

	for key, value := range m {
		if value <= max_value {
			continue
		}
		max_value = value
		max_key = key
	}
	return max_key
}

func swapMinMax(arr []int) []int {
	max_num := math.MinInt
	min_num := math.MaxInt
	var max_index, min_index int

	for index, num := range arr {
		if num > max_num {
			max_num = num
			max_index = index
		}
		if num < min_num {
			min_num = num
			min_index = index
		}
	}

	new_arr := make([]int, len(arr))
	copy(new_arr, arr)
	new_arr[max_index] = int(min_num)
	new_arr[min_index] = int(max_num)

	return new_arr
}

func swap(ptr_a, ptr_b *int) {
	ori_a := *ptr_a

	*ptr_a = *ptr_b
	*ptr_b = ori_a
}

func initValue(ptr *int) {
	*ptr = 100
}

func squareAll(ptr_nums *[]int) {
	nums := *ptr_nums
	for index, num := range nums {
		nums[index] = num * num
	}
}

func getLength(ptr_s *string) int {
	return len(*ptr_s)
}

func safeAssign(ptr_target *int, ptr_value *int) {
	if ptr_value == nil {
		return
	}
	*ptr_target = *ptr_value
}

func main() {
	fmt.Println("返回从 1 到 n 中所有偶数的和:")
	fmt.Println(sumEven(6))

	fmt.Println("接收一个整数切片并返回其反转结果:")
	fmt.Println(reverseArray([]int{1, 2, 3, 4}))

	fmt.Println("统计切片中的奇数个数:")
	fmt.Println(countOdd([]int{1, 2, 3, 4, 5}))

	fmt.Println("检查一个字符串是否为回文:")
	fmt.Println(isPalindrome("level:"))

	fmt.Println("构造乘法表:")
	printTable(3)

	fmt.Println("找出 map 中值最大的键:")
	fmt.Println(findMaxKey(map[string]int{"a": 1, "b": 10, "c": 5}))

	fmt.Println("实现一个函数，交换切片中最小值与最大值的位置:")
	fmt.Println(swapMinMax([]int{2, 9, 3, 1, 6}))

	fmt.Println("交换两个整数:")
	a, b := 3, 5
	swap(&a, &b)
	fmt.Println(a, b)

	fmt.Println("为变量赋初值（用指针）:")
	var x int
	initValue(&x)
	fmt.Println(x)

	fmt.Println("平方每个切片元素:")
	arr_a := []int{1, 2, 3}
	squareAll(&arr_a)
	fmt.Println(arr_a)

	fmt.Println("统计字符串长度:")
	str := "hello"
	length := getLength(&str)
	fmt.Println(length)

	fmt.Println("实现一个“安全赋值”函数:")
	var c int = 0
	var d *int = nil
	safeAssign(&c, d)
	fmt.Println(c)

	e := 10
	safeAssign(&c, &e)
	fmt.Println(c)
}
