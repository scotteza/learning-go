package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, everyone")

	fmt.Println("The time is", time.Now())

	fmt.Println("Here is a random number: ", rand.Intn(10))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))
}

func add(x int, y int) int {
	return x + y
}
