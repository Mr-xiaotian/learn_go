package main

import "fmt"

func chapter_2() {
	a, b := 6, 8
	fmt.Println(a+b, a-b, a*b, a/b, a%b)

	const pi float64 = 3.14159
	fmt.Println(pi * 5 * 5)

	name := "Xiaotian"
	fmt.Printf("Hello, %s\n", name)

	var test_var int
	fmt.Println(test_var)

	var (
		varInt   int     = 1
		varFloat float64 = 1.0
		varStr   string  = "test"
		varBool  bool    = true
	)
	fmt.Printf("varInt: %d, varFloat: %f, varStr: %s, varBool: %t\n", varInt, varFloat, varStr, varBool)
	fmt.Printf("varInt type: %T, varFloat type: %T, varStr type: %T, varBool type: %T\n", varInt, varFloat, varStr, varBool)
}
