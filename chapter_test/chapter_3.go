package main

import "fmt"

func chapter_3() {
	// 数组转切片与切片操作
	intList := [5]int{1, 2, 3, 4, 5}
	intSlice := make([]int, len(intList))
	copy(intSlice, intList[:])

	fmt.Println(intSlice[:3])
	fmt.Println(intSlice[2:5])
	intSlice = append(intSlice, 6, 7)
	fmt.Println(intSlice)

	// 字符串处理
	str := "Go is awesome!"
	fmt.Println(len(str))
	for index, char := range []rune(str) {
		fmt.Println(index, char)
	}
}
