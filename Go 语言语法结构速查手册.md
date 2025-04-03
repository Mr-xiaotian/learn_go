# 📘 Go 语言语法结构速查手册（完整版）

本手册覆盖 Go 语言的常见语法结构、标准写法、易错点，适合入门学习与日常查阅使用。

---

## ✅ 1. 基本变量与常量声明

| 写法                 | 说明                     |
|----------------------|--------------------------|
| `var x int`          | 声明整型变量，默认值为 0 |
| `var s string = "hi"`| 显式赋值                 |
| `x := 42`            | 简洁声明（函数内部使用） |
| `const Pi = 3.14`    | 常量声明                 |

---

## 🔣 2. 字面量（Literals）

| 类型     | 示例                       |
|----------|----------------------------|
| 整数     | `123`, `0`, `-1`           |
| 浮点数   | `3.14`, `-0.001`           |
| 字符串   | `"你好"`, `"hello"`        |
| 布尔值   | `true`, `false`            |
| 字符     | `'你'`, `'A'`              |
| 切片     | `[]int{1, 2}`              |
| 映射     | `map[string]int{"a": 1}`  |
| 结构体   | `Person{"Tom", 18}` |

---

## 🧩 3. 类型转换（Type Conversion）

| 写法                      | 说明                                |
|---------------------------|-------------------------------------|
| `float64(10)`             | int → float64                       |
| `string(65)`              | int → 字符 → "A"            |
| `[]rune("你好")`          | string → 字符切片（按 Unicode）       |
| `[]byte("abc")`          | string → 字节切片string → 字节切片（UTF-8）    |

---

## 🧰 4. make vs new（超重要！）

| 写法                         | 类型        | 说明                                 |
|------------------------------|-------------|--------------------------------------|
| `make([]int, 5)`             | 切片        | 创建并初始化切片                     |
| `make(map[string]int)`       | 映射        | 创建可用的 map                       |
| `new(int)`                   | *int（指针）       | 返回 *int，初始值为 0                |
| `new(MyStruct)`             | *MyStruct     | 返回结构体指针，初始值为零值         |

---

## 🍱 5. 切片操作
| 操作                        | 说明        |
|------------------------------|-------------|
| `a := []int{1, 2, 3}`	| 创建切片 |
| `b := a[1:3]`	| 子切片，共享内存（[1] 到 [2]） |
| `c := append(a, 4)`	| 添加元素 |
| `copy(dst, src)`	| 拷贝内容，返回复制个数 |
| `[]rune("你好")`	| 按字符分割字符串 |
| `append([]int(nil), a...)`	| 安全复制切片 |

---

## 🗺️ 6. map 映射

```go
m := map[string]int{"a": 1}
m["b"] = 2
delete(m, "a")
val, ok := m["b"]  // 查询 key 是否存在
```

---

## 🧱 7. 结构体与方法

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return "Hi, I'm " + p.Name
}

p := Person{"Tom", 18}
fmt.Println(p.Greet())
```

---

## 🔄 8. 控制结构

```go
if x > 0 {
    fmt.Println("positive")
}

for i := 0; i < 10; i++ {
    fmt.Println(i)
}

switch day {
case "Mon", "Tue":
    fmt.Println("Work")
default:
    fmt.Println("Rest")
}
```

---

## 📦 9. 函数、返回值、命名返回值

```go
func add(a, b int) int {
    return a + b
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // 返回命名返回值
}
```

---

## 💡 10. 字符串与字符处理

| 操作                   | 说明                         |
|------------------------|------------------------------|
| `s := "你好"`          | UTF-8 编码字符串（6字节）     |
| `s[0]`                 | 第一个字节，不是“你”          |
| `[]rune(s)`            | 转为字符切片（rune）          |
| `string([]rune(s)[0])` | 获取第一个字符“你”            |
| `len(s)`               | 字节长度                      |
| `len([]rune(s))`       | 字符数（真正的长度）          |

---

## 🔧 11. 指针与引用

```go
x := 10
p := &x // p 是 *int 类型
*p = 20 // 修改 x 的值
fmt.Println(x) // 20
```

---

## 🔀 12. 接口（interface）

```go
type Greeter interface {
    Greet() string
}

type Person struct {
    Name string
}

func (p Person) Greet() string {
    return "Hi, " + p.Name
}

var g Greeter = Person{"Alice"}
fmt.Println(g.Greet())
```

---

## 🚦 13. 并发（goroutine, channel）

```go
// goroutine
func say(msg string) {
    fmt.Println(msg)
}
go say("Hello")

// channel
ch := make(chan int)
go func() {
    ch <- 42
}()
fmt.Println(<-ch)  // 接收值
```

---

## 🎁 14. 常用内置函数

| 函数                  | 说明                             |
|-----------------------|----------------------------------|
| `len(s)`              | 长度（数组/切片/字符串）         |
| `cap(s)`              | 容量（切片/通道）                 |
| `append(slice, x)`    | 向切片添加元素                   |
| `copy(dst, src)`      | 拷贝切片数据                     |
| `delete(map, key)`    | 删除 map 项                      |
| `close(ch)`           | 关闭通道                         |

---

## ✅ Go 常见类型传参行为总览表

| 类型               | 是值类型？ | 默认传递行为        | 是否共享底层数据 | 能否修改原值 | 是否需要指针（`*T`） | 推荐做法                                  |
|--------------------|------------|---------------------|------------------|---------------|----------------------|-------------------------------------------|
| `int`, `float64`   | ✅ 是       | 拷贝值              | ❌ 不共享         | ❌ 否          | ✅ 需要指针才能修改   | 修改原值时传 `*int`, `*float64`           |
| `string`           | ✅ 是       | 拷贝引用（不可变）  | ✅ 共享           | ❌ 否          | ❌ 指针也没用         | 不可修改字符，建议作为值使用              |
| `array`            | ✅ 是       | 拷贝整个数组        | ❌ 不共享         | ❌ 否          | ✅ 用 `*[N]T`         | 不推荐直接用，推荐改用切片                |
| `struct`           | ✅ 是       | 拷贝整个结构体      | ❌ 不共享         | ❌ 否          | ✅ 用 `*Struct`        | ⚠️ 结构体较大或需修改字段时建议用指针     |
| `slice` (`[]T`)    | ❌ 否       | 拷贝 slice 结构体   | ✅ 共享底层数组   | ✅ 可以修改内容 | ❌ 不需要             | 直接传即可，除非要改变容量/指针行为       |
| `map[K]V`          | ❌ 否       | 拷贝 map 引用       | ✅ 共享底层数据   | ✅ 可以修改     | ❌ 不需要             | 直接传即可                                 |
| `chan T`           | ❌ 否       | 拷贝通道引用        | ✅ 同一个 channel | ✅ 可以读写     | ❌ 不需要             | 直接传即可                                 |
| `func`             | ❌ 否       | 拷贝函数引用        | ✅ 共用闭包数据   | ✅ 可调用共享逻辑 | ❌ 不需要             | 直接传即可                                 |
| `interface{}`      | ❌ 否       | 拷贝接口结构        | ✅ 指向同一数据   | ✅ 可操作共享对象 | ❌ 不需要             | 一般用于泛型场景，不用指针                 |
