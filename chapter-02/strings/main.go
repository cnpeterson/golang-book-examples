package main

import "fmt"

type Celcius float64
type Fahrenheit float64

func (c Celcius) String() string { return fmt.Sprintf("%gC", c)}

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func main() {
    c := FToC(212.0)
    fmt.Println(c.String())
    fmt.Printf("%v\n", c) // Does not call String explicitly
    fmt.Printf("%s\n", c)
    fmt.Println(c)
    fmt.Printf("%g\n", c) // Does not call string
    fmt.Println(float64(c)) // Does not call string

}
