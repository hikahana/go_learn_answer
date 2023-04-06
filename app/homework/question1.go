package main

import (
	"fmt"
	"math"
)

func sqrt (x float64) float64 {
	return math.Sqrt(x)
}

func pow (x float64, y float64) float64 {
	return math.Pow(x, y)
}

func main() {
	x := 10.0
	y := 4.0
	fmt.Println(sqrt(x))
	fmt.Println(pow(x, y))

}
