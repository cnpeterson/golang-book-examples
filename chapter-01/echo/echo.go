package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    EchoOne()
    EchoTwo()
    EchoThree()
    EchoPractice()
}

func EchoOne() {
    start := time.Now()
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s, time.Since(start))
}

func EchoTwo() {
    start := time.Now()
    var s, sep string
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s, time.Since(start))
}

func EchoThree() {
    start := time.Now()
    fmt.Println(strings.Join(os.Args[1:], " "), time.Since(start))
}

func EchoPractice() {
    //1.1
    start := time.Now()
    fmt.Println(strings.Join(os.Args[0:], " "), time.Since(start))

    //1.2
    start = time.Now()
    for i, arg := range os.Args[1:] {
        fmt.Println(i + 1, arg)
    }

    fmt.Println(time.Since(start))
}

