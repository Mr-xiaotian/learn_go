package main

import "math"

func reverseArray(arr []int) []int {
	reverse_array := []int{}
	arr_max_index := len(arr) - 1

	for index := range arr {
		reverse_array = append(reverse_array, arr[arr_max_index-index])
	}
	return reverse_array
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
	new_arr[max_index] = min_num
	new_arr[min_index] = max_num

	return new_arr
}

func squareAll(ptr_nums *[]int) {
	nums := *ptr_nums
	for index, num := range nums {
		nums[index] = num * num
	}
}

func appendToSlice(s *[]int, val int) {
	*s = append(*s, val)
}
