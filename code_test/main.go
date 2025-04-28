// 📌 main.go
package main

import "fmt"

func main() {
	fmt.Println("🧪 Go Test Hub - Select the tests you want to run")

	// Uncomment the tests you want to execute:
	// testMath()
	// testSlice()
	// testString()
	// testMap()
	// testPointer()
	// testStruct()
	// testFactory()
	testInterface()
}

// ---------------------
// 🔢 Math Tests
func testMath() {
	fmt.Println("\n🔢 Math Tests:")
	fmt.Println("sumEven(6) =", sumEven(6))
	fmt.Println("countOdd([1,2,3,4,5]) =", countOdd([]int{1, 2, 3, 4, 5}))
	printTable(3)
}

// 🔁 Slice Tests
func testSlice() {
	fmt.Println("\n🔁 Slice Tests:")
	fmt.Println("reverseArray([1,2,3,4]) =", reverseArray([]int{1, 2, 3, 4}))
	fmt.Println("swapMinMax([2,9,3,1,6]) =", swapMinMax([]int{2, 9, 3, 1, 6}))
	arr := []int{1, 2, 3}
	squareAll(&arr)
	fmt.Println("squared arr =", arr)
	s := []int{1, 2}
	appendToSlice(&s, 99)
	fmt.Println("appended slice =", s)
}

// 🔤 String Tests
func testString() {
	fmt.Println("\n🔤 String Tests:")
	fmt.Println("reverseString('hello') =", reverseString("hello"))
	fmt.Println("isPalindrome('level') =", isPalindrome("level"))
	str := "hello"
	fmt.Println("getLength('hello') =", getLength(&str))
}

// 🗺️ Map Tests
func testMap() {
	fmt.Println("\n🗺️ Map Tests:")
	m := map[string]int{"a": 1, "b": 10, "c": 5}
	fmt.Println("findMaxKey(m) =", findMaxKey(m))
}

// 🧷 Pointer Tests
func testPointer() {
	fmt.Println("\n🧷 Pointer Tests:")
	a, b := 3, 5
	swap(&a, &b)
	fmt.Println("swapped a, b =", a, b)

	var x int
	initValue(&x)
	fmt.Println("initValue =", x)

	var c int = 0
	var d *int = nil
	safeAssign(&c, d)
	e := 10
	safeAssign(&c, &e)
	fmt.Println("safe assigned c =", c)

	x1, y1 := 3, 7
	pa, pb := &x1, &y1
	swapDoublePointer(&pa, &pb)
	fmt.Println("double swap =", *pa, *pb)

	x_ptr := createPtr()
	fmt.Println("createPtr result =", *x_ptr)
}

// 👥 Struct Tests
func testStruct() {
	fmt.Println("\n👥 Struct Tests:")
	add := Address{"SuZhou"}
	user := User{"V", &add}
	copy_user := cloneUser(&user, "Tokyo")
	fmt.Println("copied user =", *copy_user)
	fmt.Println("copied address =", *copy_user.Addr)
	circle := Circle{4}
	fmt.Println("Circle Area =", circle.Area())
	person := Person{"Lion", 18}
	person.HaveBirthday()
	fmt.Println("person age =", person.Age)
	student_0 := Student{"Alice", 15}
	student_1 := Student{"Bob", 21}
	student_2 := Student{"Carl", 12}
	fmt.Println("average age =", AverageAge([]Student{student_0, student_1, student_2}))
	author := Author{"小明"}
	book := Book{"Go语言之旅", author}
	fmt.Println("book info =", book.GetInfo())
	rectangle := Rectangle{20, 19}
	fmt.Println("new rectangle = ", rectangle.Scale(1.5))
}

// 🎛️ Function Factory Test
func testFactory() {
	fmt.Println("\n🎛️ Function Factory Test:")
	inc := MakeIncrementor(5)
	fmt.Println(*inc())
	fmt.Println(*inc())
	fmt.Println(*inc())
}

func testInterface() {
	fmt.Println("\nInterface Test:")
	cicle := NewCircle{5}
	rectangel := NewRectangle{2, 8}
	shapeslice := []Shape{cicle, rectangel}
	for _, s := range shapeslice {
		PrintArea(s)
		checkType(s)
	}

	consolewriter := ConsoleWriter{}
	consolewriter.Write("It's console write")
}
