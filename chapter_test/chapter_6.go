package main

import "fmt"

func swapPtr(a, b *int) {
	*a, *b = *b, *a
}

type Person struct {
	Name string
	Age  int
}

func growUp(p *Person) {
	p.Age += 1
}

func printPtr(num *int) {
	if num == nil {
		fmt.Println("空指针")
	} else {
		fmt.Println(*num)
	}
}

func chapter_6() {
	// 交换函数升级版（指针）
	a := 15
	b := 22
	swapPtr(&a, &b)
	fmt.Println(a, b)

	// 结构体指针操作
	person := Person{"Wang", 33}
	growUp(&person)
	fmt.Println(person)

	// 创建一个 int 指针并赋值
	intPtr := new(int)
	fmt.Printf("Int value: %d\nInt address: %p\nPtr address: %p\n", *intPtr, intPtr, &intPtr)

	// 判断指针是否为空
	printPtr(new(int))
}
