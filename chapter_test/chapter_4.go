package main

import "fmt"

func shadow_var() {
	x := 10
	if true {
		fmt.Println(x)
		x := 20
		fmt.Println(x)
	}
	fmt.Println(x)
}

func score_switch(score int) {
	switch true {
	case 90 <= score && score <= 100:
		fmt.Println("优秀")
	case 70 <= score && score <= 89:
		fmt.Println("良好")
	case 60 <= score && score <= 69:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

func infinite_loop() {
	index := 1
	for {
		if index >= 5 {
			break
		}
		fmt.Println(index)
		index += 1
	}
}

func chapter_4() {
	// 影子变量演示程序
	shadow_var()

	// 九九乘法表输出
	for j := range 10 {
		for i := range j {
			fmt.Printf("%d * %d = %d, ", i+1, j, (i+1)*j)
		}
		fmt.Println()
	}

	// switch分支匹配演示器
	score_switch(75)

	// 无限循环与 break 控制
	infinite_loop()
}
