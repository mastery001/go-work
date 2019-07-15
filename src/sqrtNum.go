package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 1 ; i < 10 ; i++  {
		z -= (z * z - x) / (2 * z)
		fmt.Printf("loop %d res is %p\n" , i , z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
}