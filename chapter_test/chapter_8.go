package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("errors: divisor isn't 0")
	}
	return a / b, nil
}

func divid_with_errorf(a, b int) (int, error) {
	quotient, err := divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("divide failed: %w", err)
	}
	return quotient, nil
}

type NegativeError struct {
	Value int
}

func (n NegativeError) Error() string {
	return fmt.Sprintf("negative value: %d", n.Value)
}

func test_NegativeError(num int) (int, error) {
	if num < 0 {
		return 0, NegativeError{num}
	}
	return num, nil
}

func safeDivide(a, b int) int {
	if b == 0 {
		panic("panic: divisor isn't 0")
	}
	return a / b
}

func chapter_8() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	// 标准错误返回
	fmt.Println(divide(10, 5))
	fmt.Println(divide(3, 0))

	// 使用 fmt.Errorf 包装错误
	_, err := divid_with_errorf(9, 0)
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	// 自定义错误类型
	fmt.Println(test_NegativeError(5))
	fmt.Println(test_NegativeError(-5))

	// panic / recover 捕获演示
	safeDivide(2, 0)
}
