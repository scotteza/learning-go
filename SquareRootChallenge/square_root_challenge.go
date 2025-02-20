package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	previousZ := 0.0
	for i := 0; i < 10000; i++ {
		z -= (z*z - x) / (2 * z)
		if i == 0 {
			previousZ = z
		} else {
			if math.Abs(z-previousZ) < 0.0000000000000001 {
				break
			}
		}
	}
	return z
}

func main() {
	x := 20000000.0
	calculatedSquareRoot := Sqrt(x)
	mathLibrarySquareRoot := math.Sqrt(x)

	fmt.Printf("Calculated square root is %v\n", calculatedSquareRoot)
	fmt.Printf("Math library square root is %v\n", mathLibrarySquareRoot)
	fmt.Printf("Difference between them is %v\n", math.Abs(calculatedSquareRoot-mathLibrarySquareRoot))
}
