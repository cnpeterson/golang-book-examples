package main

import (
    "fmt"
    "os"
    "strconv"
)

type Celcius float64
type Fahrenheit float64
type Feet float64
type Meters float64
type Pounds float64
type Kilograms float64

func main() {
    argsWithoutProg := os.Args[1:]
    if len(argsWithoutProg) == 0 {
        fmt.Println("Enter a number you want converted: ")
        var a string
        fmt.Scanln(&a)
        argsWithoutProg = append(argsWithoutProg, a)
    }
    for _, arg := range argsWithoutProg {
        i, err := strconv.Atoi(arg)
        if err != nil {
            fmt.Printf("%s is not a number skipping\n", arg)
            continue
        }

        fmt.Printf("**** Number to be converted: %d ****\n", i)
        fmt.Printf("Celcius: %dC == Fahrenheit: %gF\n", i, CToF(Celcius(i)))
        fmt.Printf("Fahrenheit: %dF == Celcius: %gC\n", i, FToC(Fahrenheit(i)))
        fmt.Printf("Feet: %d\" == Meters: %gm\n", i, FToM(Feet(i)))
        fmt.Printf("Meters: %dm == Feet: %g\"\n", i, MToF(Meters(i)))
        fmt.Printf("Pounds: %dlb == Kilograms: %gkg\n", i, PToK(Pounds(i)))
        fmt.Printf("Kilograms: %dkg == Pounds: %glb\n", i, KToP(Kilograms(i)))
    }
}


func CToF(c Celcius) Fahrenheit { return Fahrenheit(c * 9/5 + 32) }

func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5/9) }

func FToM(fe Feet) Meters { return Meters(fe / 3.281) }

func MToF(me Meters) Feet { return Feet(me * 3.281) }

func PToK(p Pounds) Kilograms { return Kilograms(p / 2.205) }

func KToP(k Kilograms) Pounds { return Pounds(k * 2.205) }
