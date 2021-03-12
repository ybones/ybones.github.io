package main

import (
    "fmt"
    "unsafe"
)
// bool sizeof  1
// int sizeof  8
// int sizeof  8
// int8 sizeof  1
// int16 sizeof  2
// int32 sizeof  4
// int64 sizeof  8
// uint sizeof  8
// uint8 sizeof  1
// uint16 sizeof  2
// uint32 sizeof  4
// uint64 sizeof  8
// array sizeof  40     与数组长度和数组类型有关
// slice 1 sizeof  24
// slice 2 sizeof  24
// unsafe.Pointer 1 sizeof  8
// unsafe.Pointer 2 sizeof  8
// uintptr sizeof  8
// uintptr sizeof  8
// chan sizeof  8
// map sizeof  8
// func sizeof 8
// interface sizeof 16
// string sizeof 16
// struct sizeof 40    字节对齐
func main() {
    boolT := false
    fmt.Println("bool sizeof ", unsafe.Sizeof(boolT))
    intTT := 10
    fmt.Println("int sizeof ", unsafe.Sizeof(intTT))
    var intT int = 10
    var int8T int8 = 10
    var int16T int16 = 10
    var int32T int32 = 10
    var int64T int64 = 10
    fmt.Println("int sizeof ", unsafe.Sizeof(intT))
    fmt.Println("int8 sizeof ", unsafe.Sizeof(int8T))
    fmt.Println("int16 sizeof ", unsafe.Sizeof(int16T))
    fmt.Println("int32 sizeof ", unsafe.Sizeof(int32T))
    fmt.Println("int64 sizeof ", unsafe.Sizeof(int64T))
    var uintT uint = 10
    var uint8T uint8 = 10
    var uint16T uint16 = 10
    var uint32T uint32 = 10
    var uint64T uint64 = 10
    fmt.Println("uint sizeof ", unsafe.Sizeof(uintT))
    fmt.Println("uint8 sizeof ", unsafe.Sizeof(uint8T))
    fmt.Println("uint16 sizeof ", unsafe.Sizeof(uint16T))
    fmt.Println("uint32 sizeof ", unsafe.Sizeof(uint32T))
    fmt.Println("uint64 sizeof ", unsafe.Sizeof(uint64T))

    // 数组
    arrayT := [5]int64{1, 2}
    fmt.Println("array sizeof ", unsafe.Sizeof(arrayT))
    // 切片
    sliceT := arrayT[1:]
    fmt.Println("slice 1 sizeof ", unsafe.Sizeof(sliceT))
    sliceT2 := make([]int64, 0)
    fmt.Println("slice 2 sizeof ", unsafe.Sizeof(sliceT2))
    // unsafe.Pointer
    point8T := unsafe.Pointer(&int8T)
    fmt.Println("unsafe.Pointer 1 sizeof ", unsafe.Sizeof(point8T))
    pointT := unsafe.Pointer(&intT)
    fmt.Println("unsafe.Pointer 2 sizeof ", unsafe.Sizeof(pointT))

    // uintptr
    var uintptrT uintptr = 10
    fmt.Println("uintptr sizeof ", unsafe.Sizeof(uintptrT))

    uintptrT2 := (uintptr)(unsafe.Pointer(&uintT))
    fmt.Println("uintptr sizeof ", unsafe.Sizeof(uintptrT2))

    // chan
    chanT := make(chan int, 1)
    fmt.Println("chan sizeof ", unsafe.Sizeof(chanT))

    // map
    mapT := make(map[int]string)
    fmt.Println("map sizeof ", unsafe.Sizeof(mapT))

    // func
    funcT := func() {}
    fmt.Println("func sizeof", unsafe.Sizeof(funcT))

    // interface
    type interfaceT interface {
    }
    var interT interfaceT
    fmt.Println("interface sizeof", unsafe.Sizeof(interT))

    // string
    strT := "hello"
    fmt.Println("string sizeof", unsafe.Sizeof(strT))

    // struct
    type StructTemp struct {
        Field3 bool   // 1
        Field2 int    //8
        Field4 uint64 // 8
        Field1 string //16
    }
    structT := StructTemp{true, 1, 10, "zhangsan"}
    fmt.Println("struct sizeof", unsafe.Sizeof(structT))
}
