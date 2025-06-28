package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var i int
    var f float64
    var b bool
    var s string
    
    fmt.Printf("int: %d bytes\n", unsafe.Sizeof(i))
    fmt.Printf("float64: %d bytes\n", unsafe.Sizeof(f))
    fmt.Printf("bool: %d bytes\n", unsafe.Sizeof(b))
    fmt.Printf("string: %d bytes\n", unsafe.Sizeof(s))
}