package main

import (
	"fmt"
	"math"
)

// 🧠 Go 接口基础题目（适合入门）
// 🔹题目 1：定义并实现一个简单接口
// 定义一个接口 Shape，包含一个方法：
// Area() float64
// 实现两个结构体 Circle 和 Rectangle，使它们满足 Shape 接口。

type Shape interface {
	Area() float64
}

type NewCircle struct {
	Radius float64
}

func (c NewCircle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type NewRectangle struct {
	Length float64
	Width  float64
}

func (r NewRectangle) Area() float64 {
	return r.Length * r.Width
}

// 🔹题目 2：接口作为函数参数
// 定义如下函数：
// func PrintArea(s Shape)
// 它接受任意 Shape 类型，输出其面积。用 Circle 和 Rectangle 测试它。

func PrintArea(s Shape) {
	fmt.Print(s.Area(), " ")
}

// 🔹题目 3：实现一个接口类型的切片
// 创建一个 []Shape 类型的切片，里面放入多个不同类型的图形（如 Circle, Rectangle），遍历这个切片，输出所有图形的面积。

// 🔹题目 4：使用类型断言获取原始类型
// 继续用上题中的 []Shape，增加逻辑：
// 如果某个 Shape 是 Circle，则打印 "It's a circle"
// 如果是 Rectangle，打印 "It's a rectangle"
// （使用 type switch 或 类型断言 实现）

func checkType(s Shape) {
	switch s.(type) {
	case NewCircle:
		fmt.Println("Wow! It's a circle!")
	case NewRectangle:
		fmt.Println("Hoo! It's a rectangle!")
	}
}

// 🔹题目 5：实现一个 Writer 接口模拟
// 定义一个接口：
//
//	type Writer interface {
//	    Write(content string)
//	}
//
// 再定义一个结构体 ConsoleWriter，让它实现 Writer 接口，并在 Write() 方法中打印内容到控制台。
type Writer interface {
	Write(content string)
}

type ConsoleWriter struct{}

func (w ConsoleWriter) Write(content string) {
	fmt.Println(content)
}
