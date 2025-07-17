# Go语言学习指南 - 第十章复习题

📘 主题：goroutine、channel、select、mutex、并发安全

---

## 一、选择题（每题1分）

1. 在 Go 中创建新的 goroutine 的语法是：

   * A. `start function()`
   * B. `fork function()`
   * C. `go function()`
   * D. `thread function()`

2. 默认的 channel 是否带缓冲？

   * A. 是
   * B. 否
   * C. 取决于变量名
   * D. 被系统自动决定

3. 下列关于 `select` 的说法正确的是：

   * A. 它可以同时监听多个 channel
   * B. 它只能监听一个 channel
   * C. 它只能用于 buffered channel
   * D. 它用于处理 map 的并发读写

4. 下列哪一项是互斥锁的标准使用方式？

   * A. `lock()` / `unlock()`
   * B. `mutex.Lock()` / `mutex.Unlock()`
   * C. `sync(lock)` / `sync(unlock)`
   * D. `begin / end mutex`

---

## 二、填空题（每空1分）

5. goroutine 是由 Go 运行时调度的轻量级 \_\_\_\_\_\_\_。

6. channel 是 Go 中用于 \_\_\_\_\_\_\_ 和 \_\_\_\_\_\_\_ 的通信机制。

7. 使用 `go` 启动新 goroutine 后，若主线程先退出，子 goroutine 将 \_\_\_\_\_\_\_。

8. 使用 `sync.WaitGroup` 时，常用的三个方法是：\_\_\_\_\_\_\_、\_\_\_\_\_\_\_、\_\_\_\_\_\_\_。

---

## 三、简答题（每题4分）

9. 简述 channel 和互斥锁的区别，各自适合什么场景？是否可以组合使用？

10. select 的作用是什么？当所有 channel 都无法使用时会发生什么？如何处理这个情况？

---

## 四、编程题（共4题，每题6分，共24分）

11. **并发打印数字**
    启动 5 个 goroutine，每个 goroutine 打印从 1 到 5 的数字（每个 goroutine 标注编号），主函数使用 `WaitGroup` 等待全部完成。

12. **通道通信**
    创建一个 channel，主 goroutine 向其中发送 10 个整数，另一个 goroutine 从中接收并打印。使用无缓冲 channel。

13. **select 多路监听**
    使用两个 goroutine 分别向两个 channel 发送数据，主 goroutine 使用 `select` 同时监听并打印接收到的值。

14. **互斥锁保证并发安全**
    创建一个共享整数变量，启动 100 个 goroutine，每个将其加 1。使用 `sync.Mutex` 保证加法操作的并发安全，并输出最终值。
