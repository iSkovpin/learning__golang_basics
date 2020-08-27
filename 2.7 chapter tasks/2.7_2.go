package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(calcHypotenuse(a, b))
}

func calcHypotenuse(leg1 float64, leg2 float64) float64 {
	return math.Sqrt(leg1*leg1 + leg2*leg2)
}
