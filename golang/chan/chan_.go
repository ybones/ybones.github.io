package main

import (
    "fmt"
    "time"
)

// 只读chan
func ReadChan(c <-chan int) {
    i := <-c
    fmt.Println("ReadChan", i)
}

// 只写chan
func WriteChan(c chan<- int, v int) {
    c <- v
}

var count int32 = 0

// 使用chan实现简易锁
func SyncTest(c chan int) {
    c <- 1
    count++
    fmt.Println(count)
    <-c
}

// 无缓冲的chan
func EmptyChanTest() {
    c := make(chan int)
    <-c    // deadlock!
    c <- 1 // deadlock!
}

// 关闭chan测试
// 已关闭的chan，进行写入会panic
// 已关闭的chan，有数据可以继续读，没有则会阻塞
func CloseChanTest() {
    c := make(chan int, 3)
    c <- 1
    c <- 2
    close(c)
    // c <- 3 //panic
    <-c
    <-c
    fmt.Println("hello")
}

// 关闭chan测试2
func CloseChanTest2() {
    c := make(chan int, 3)
    c <- 1
    c <- 2
    close(c)
    v, ok := <-c
    fmt.Println("v", v, "ok", ok) // v 1 ok true
    v, ok = <-c
    fmt.Println("v", v, "ok", ok) // v 2 ok true
    v, ok = <-c
    fmt.Println("v", v, "ok", ok) // v 0 ok false
}

// 关闭值为nil的chan
func CloseChanTest3() {
    var c chan int
    fmt.Println("c", c) // nil
    close(c)            // panic
}

// 关闭 已经关闭的chan
func CloseChanTest4() {
    c := make(chan int, 1)
    close(c)
    close(c) // panic
}

// 循环读取chan
// 系统检测无写入数据的goroutine，会deadlock!
func RangeChanTest(c chan int) {
    for _c := range c {
        fmt.Println("_c", _c)
    }
}

func main() {
    c := make(chan int, 1)
    go RangeChanTest(c)
    go func() {
        for{
            c<-1
            time.Sleep(time.Second*2)
        }
    }()
    time.Sleep(time.Second * 10)
}
