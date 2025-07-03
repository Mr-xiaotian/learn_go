package main

import (
	"fmt"
	"math"
)

type MyInt int

func (i MyInt) Double() MyInt {
	return i * 2
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (dog Dog) Speak() string {
	return fmt.Sprintln("汪汪")
}

func describe(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("字符串")
	case int:
		fmt.Println("整数")
	case MyInt:
		fmt.Println("自定义整数")
	default:
		fmt.Println("未知类型")
	}
}

type Shaper interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea(s Shaper) {
	fmt.Println(s.Area())
}

func chapter_7() {
	// 方法绑定演示
	var i MyInt = 9
	fmt.Printf("i: %d, i type: %T\n", i, i)
	new_i := i.Double()
	fmt.Printf("i: %d, i type: %T\n", new_i, new_i)

	// 结构体与方法演示
	r := Rectangle{50, 7}
	fmt.Printf("Area: %f\n", r.Area())

	// 接口实现演示
	dog := Dog{"Joke"}
	fmt.Print(dog.Speak())

	// 空接口与类型断言
	describe(5)
	describe(MyInt(10))
	describe("16")
	describe('s')

	// 多态函数
	c := Circle{5}
	printArea(c)
	printArea(r)
}
