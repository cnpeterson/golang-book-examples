package main

import (
    "cnpeterson/golang-book-examples/chapter-02/popcount"
    "fmt"
    "bytes"
    "time"
)

var h = map[int]int{}

func main() {
    start := time.Now()
    for n := 0; n < 26; n++ {
        x := popcount.PopCount(uint64(n))
        h[n] = x
    }
    fmt.Println(time.Since(start))
    formatOutput(h)

    h = map[int]int{}
    start = time.Now()
    for n := 0; n < 26; n++ {
        x := popcount.PopCountPerf(uint64(n))
        h[n] = x
    }
    fmt.Println(time.Since(start))
    formatOutput(h)

    h = map[int]int{}
    start = time.Now()
    for n := 0; n < 26; n++ {
        x := popcount.PopCountShifty(uint64(n))
        h[n] = x
    }
    fmt.Println(time.Since(start))
    formatOutput(h)

    h = map[int]int{}
    start = time.Now()
    for n := 0; n < 26; n++ {
        x := popcount.PopCountClearing(uint64(n))
        h[n] = x
    }
    fmt.Println(time.Since(start))
    formatOutput(h)
}

func formatOutput(a map[int]int) {
    var top bytes.Buffer
    var bottom bytes.Buffer
    writeStringForMe("|", &top)
    writeStringForMe("|", &bottom)
    alen := len(a)
    for key := 0; key < alen; key++ {
        value := a[key]
        x := fmt.Sprintf(" %d ", key)
        writeStringForMe(x, &top)
        x = fmt.Sprintf(" %d ", value)
        writeStringForMe(x, &bottom)
        rdc := recursiveDigitCounter(key)
        for i := 1; i < rdc; i++ {
            writeStringForMe(" ", &bottom)
        }
        writeStringForMe("|", &top)
        writeStringForMe("|", &bottom)
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

func writeStringForMe(s string, w* bytes.Buffer) {
    w.WriteString(s)
}
