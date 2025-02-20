package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	c, python, java bool
	i, j            int  = 1, 2
	x1, x2               = 3, false
	check_my_type   bool = false
)

func main() {
	fmt.Println("Hello, everyone")

	fmt.Println("The time is", time.Now())

	fmt.Println("Here is a random number: ", rand.Intn(10))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	fmt.Println(add_many(1, 2, 3, 4, 5, 6))

	fmt.Println(add_two(1, 2, 3))

	fmt.Println(split(17))

	x, y := split(17)

	fmt.Println(x)
	fmt.Println(y)

	k := 3

	// Note how this overrides the value of I from the outer scope
	var i int = 42
	fmt.Println("i is ", i)
	fmt.Println(i, c, python, java, i, j, x1, x2, k)

	fmt.Printf("Type: %T, Value: %v\n", check_my_type, check_my_type)
}

func add(x int, y int) int {
	return x + y
}

func add_many(x, y, z, a, b, c int) int {
	return x + y + z + a + b + c
}

func add_two(x, y, z int) (int, int, int, string) {
	return x + 2, y + 2, z + 2, "Finished adding 2"
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
