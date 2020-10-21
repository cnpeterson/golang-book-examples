package main

import (
    "bufio"
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)

func main() {
    ReadFromFiles()
}

func ReadFromStdin() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func ReadFromFiles() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        CountLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
                continue
            }
            CountLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)

        }
    }
}

func CountLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        fo := []string{f.Name(), input.Text()}
        s := strings.Join(fo, ": ")
        counts[s]++
    }
}

func ReadFileToMemory() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
