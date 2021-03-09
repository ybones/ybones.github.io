package main

import "fmt"

// 扩容测试
func SliceCap() {
    s := make([]int, 2)
    s = append(s, 3, 4, 5)
    fmt.Println("sLen", len(s), "sCap", cap(s)) // sLen 5 sCap 6
}

// 扩容测试2
func SliceCap1() {
    s := make([]int, 2)
    s = append(s, 3, 4, 5, 6, 7)
    fmt.Println("sLen", len(s), "sCap", cap(s)) // sLen 7 sCap 8
}

func SliceTest()  {
    var s []int
    fmt.Println(s == nil)
    fmt.Printf("s %p \n", &s)
    a := make([]int, 0)
    fmt.Println(a == nil)
    fmt.Printf("a %p \n", &a)
    b := make([]int, 0)
    fmt.Printf("b %p \n", &b)



}

func main() {
    SliceTest()
}
