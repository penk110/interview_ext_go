package main

import (
    "fmt"
    "unsafe"
)

type T struct {
    A int64
    B int64
    C int32 // 4
}

func main() {
    t := T{
        A: 1,
        B: 2,
        C: 3,
    }

    bPtr := (*[8]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&t)) + 16))

    fmt.Println(*bPtr)
    fmt.Printf("%v %b \n", *bPtr, *bPtr)
}

