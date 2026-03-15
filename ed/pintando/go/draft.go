package main

import (
    "fmt"
    "math"
)

func main() {
	var a float64
	var b float64
	var c float64
    // var p float64
    // var area float64


    _, err := fmt.Scanln(&a)
    if err != nil {
        fmt.Println("Error")

    }


    _, err = fmt.Scanln(&b)
    if err != nil {
        fmt.Println("Error")

    }


    _, err = fmt.Scanln(&c)
    if err != nil {
        fmt.Println("Error")
    }


    // p = (a+b+c)/2

    // area = p * (p - a) * (p - b) * (p - c)

    
    // fmt.Printf("%.2f\n", math.Sqrt(p * (p - a) * (p - b) * (p - c)))
    fmt.Printf("%.2f\n", math.Sqrt(((a+b+c)/2) * (((a+b+c)/2) - a) * (((a+b+c)/2) - b) * (((a+b+c)/2) - c)))
}
