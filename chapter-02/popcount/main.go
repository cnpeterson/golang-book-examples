package main

import (
    "cnpeterson/golang-book-examples/chapter-02/popcount"
    "fmt"
    "bytes"
)

var h = map[int]int{}

func main() {
    for n := 0; n < 26; n++ {
        x := popcount.PopCount(uint64(n))
        h[n] = x
    }
    formatOutput(h)
}

func formatOutput(a map[int]int) {
    var top bytes.Buffer
    var bottom bytes.Buffer
    byteWriter("|", &top)
    byteWriter("|", &bottom)
    alen := len(a)
    for key := 0; key < alen; key++ {
        value := a[key]
        x := fmt.Sprintf(" %d ", key)
        byteWriter(x, &top)
        x = fmt.Sprintf(" %d ", value)
        byteWriter(x, &bottom)
        rdc := recursiveDigitCounter(key)
        for i := 1; i < rdc; i++ {
            byteWriter(" ", &bottom)
        }
        byteWriter("|", &top)
        byteWriter("|", &bottom)
    }
    fmt.Println(top.String())
    fmt.Println(bottom.String())
}

func recursiveDigitCounter(num int) int {
    if num < 10 {
        return 1
    }else {
        return 1 + recursiveDigitCounter(num/10)
    }
}

func byteWriter(s string, w* bytes.Buffer) {
    w.WriteString(s)
}
