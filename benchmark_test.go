package main

import (
	"testing"
)

type BigStruct struct {
	Data [1024 * 1024]byte // 1MB 数据
}

func handleByValue(b BigStruct) {
	// 模拟轻量操作
	b.Data[0]++
}

func handleByPointer(b *BigStruct) {
	// 模拟轻量操作
	b.Data[0]++
}

func BenchmarkByValue(b *testing.B) {
	var s BigStruct
	for i := 0; i < b.N; i++ {
		handleByValue(s)
	}
}

func BenchmarkByPointer(b *testing.B) {
	var s BigStruct
	for i := 0; i < b.N; i++ {
		handleByPointer(&s)
	}
}

func TestDummy(t *testing.T) {
	// 只是为了测试输出，不做实际操作
}
