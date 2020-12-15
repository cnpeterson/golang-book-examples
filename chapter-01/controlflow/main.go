package main

import (
    "fmt"
    "math/rand"
    "time"
)

var heads int
var tails int

func main() {
    switch coinflip() {
    case "heads":
        heads++
        fmt.Println("Heads")
    case "tails":
        tails++
        fmt.Println("Tails")
    default:
        fmt.Println("Landed on edge")
    }
}

func coinflip() (string) {
    rand.Seed(time.Now().UTC().Unix())
    x := rand.Int() % 2
    switch x {
    case 0:
        return "heads"
    case 1:
        return "tails"
    default:
        return ""
    }
}
