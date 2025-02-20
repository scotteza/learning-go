package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)

	x := 42
	y := 3.142
	z := 0.867 + 0.5i // complex128

	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)

	const who string = "World"
	fmt.Println("Hello", who)

	fmt.Println(needInt(Small))
	// This overflows
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
