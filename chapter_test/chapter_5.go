package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func makeCounter() func() int {
	count_num := 0
	return func() int {
		count_num += 1
		return count_num
	}
}

func countVowels(s string) int {
	vowelSlice := []rune("aeiou")
	count_num := 0
	for _, char := range s {
		for _, v := range vowelSlice {
			if char != v {
				continue
			}
			count_num += 1
		}
	}
	return count_num
}

// 优化建议：使用 map[rune]bool 来提升查询效率
// func countVowels(s string) int {
// 	vowelMap := map[rune]bool{
// 		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
// 	}
// 	count := 0
// 	for _, char := range s {
// 		if vowelMap[char] {
// 			count++
// 		}
// 	}
// 	return count
// }

func deferDemo() {
	defer fmt.Println("The 1st defer")
	defer fmt.Println("The 2nd defer")
	defer fmt.Println("The 3rd defer")
}

func apply(f func(int) int, nums []int) []int {
	newSlice := make([]int, 0)
	for _, num := range nums {
		newSlice = append(newSlice, f(num))
	}
	return newSlice
}

func add_1(num int) int {
	return num + 1
}

func chapter_5() {
	// 交换函数
	fmt.Println(swap(5, 8))

	// 计数器闭包
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// 统计字符串中的元音字母数量
	fmt.Println(countVowels("werguihliuhfdsfuighergniubaerbeb"))

	// 使用 defer 打印执行顺序
	deferDemo()

	// 函数作为参数的使用
	fmt.Println(apply(add_1, []int{8, 5, 2, 3, 5, 4, 6, 9}))
}
