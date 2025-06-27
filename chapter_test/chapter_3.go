package main

import (
	"fmt"
	"sort"
)

type Address struct {
	City    string
	ZipCode int
}
type User struct {
	Name    string
	Age     int
	Address Address
}

func updateCity(u *User, newCity string) {
	u.Address.City = newCity
}

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
	for index, char := range str {
		fmt.Printf("%d: %c\n", index, char)
	}

	// map词频统计器
	randomSlice := []string{"go", "java", "go", "python", "go", "java"}
	m := map[string]int{}
	for _, v := range randomSlice {
		m[v] = m[v] + 1
	}
	fmt.Println(m)

	// 结构体排序
	type Person struct {
		Name string
		Age  int
	}
	person_0 := Person{"Anna", 23}
	person_1 := Person{"Bob", 16}
	person_2 := Person{"Angel", 888}
	personSlice := []Person{person_0, person_1, person_2}
	sort.Slice(personSlice, func(i, j int) bool {
		return personSlice[i].Age < personSlice[j].Age
	})
	fmt.Println(personSlice)

	// 嵌套结构体处理器
	add := Address{"Suzhou", 622395}
	user := User{"Crumber", 34, add}
	updateCity(&user, "Shanghai")
	fmt.Println(user)
}
