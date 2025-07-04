# Go语言学习指南 - 第八章复习题

📘 主题：错误处理、panic、recover、自定义错误

---

## 一、选择题（每题1分）

1. 在 Go 中，内置的错误接口类型是：

   * A. `Exception`
   * B. `Throwable`
   * C. `error`
   * D. `Err`

2. 使用 `errors.New("msg")` 创建的错误对象，其底层实现是：

   * A. 字符串常量
   * B. 自定义 struct
   * C. 实现了 `error` 接口的值
   * D. `panic` 类型

3. 以下关于 `panic` 的说法正确的是：

   * A. 一旦发生 panic，程序会立即退出
   * B. `panic` 可被 `recover` 捕获，用于挽救程序
   * C. panic 发生时不会执行 defer
   * D. panic 是 Go 的推荐错误处理方式

4. 在 `fmt.Errorf("wrap: %w", err)` 中 `%w` 的作用是：

   * A. 打印错误的详细信息
   * B. 忽略原始错误
   * C. 将原始错误封装以支持错误链
   * D. 转换错误为 warning

---

## 二、填空题（每空1分）

5. `error` 是一个内置接口，定义为 `___________`。

6. 若函数返回多个值，其中之一为 error 类型，通常的处理习惯是用 \_\_\_\_\_\_\_ 判断是否出错。

7. `panic` 触发后，程序会立即停止执行当前函数，并按 \_\_\_\_\_\_\_ 顺序执行已注册的 `defer` 函数。

8. 使用 `recover()` 可以在 `_______` 函数内部捕获 panic，从而避免程序崩溃。

---

## 三、简答题（每题4分）

9. Go 为什么采用显式的 error 返回值而不是 try-catch？这样设计的优劣在哪里？

10. 请说明 `panic` 和 `error` 的区别与适用场景。什么时候应该用 panic，什么时候用 error？

---

## 四、编程题（共4题，每题6分，共24分）

11. **标准错误返回**
    编写函数 `divide(a, b int) (int, error)`，当 b 为 0 时返回一个错误，否则返回 a / b。

12. **使用 fmt.Errorf 包装错误**
    在 `divide` 函数中用 `fmt.Errorf("divide failed: %w", err)` 包装原始错误并返回，调用方用 `errors.Unwrap` 展示底层错误。

13. **自定义错误类型**
    定义一个自定义错误 `NegativeError`，在传入负数时触发，包含字段 `Value int`，错误信息形如 `"negative value: -5"`。

14. **panic / recover 捕获演示**
    写一个函数 `safeDivide(a, b int) int`，当 b 为 0 时 `panic`，但在 `main` 中用 `recover` 捕获 panic 并输出错误信息而不中止程序。
