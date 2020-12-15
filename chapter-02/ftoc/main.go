package main

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celcius = -273.15
    FreezingC     Celcius = 0
    BoilingC      Celcius = 100
)

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func main() {
    const freezingF, boilingF = 32.0, 212.0
    fmt.Printf("%gF = %gC\n", freezingF, FToC(freezingF))
    fmt.Printf("%gF = %gC\n", boilingF, FToC(boilingF))
}
