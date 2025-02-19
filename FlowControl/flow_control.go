package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Infinite loop
	for {
		sum++
		if sum%100000000 == 0 {
			fmt.Println(sum)
		} else {
			fmt.Println("¯\\_(ツ)_/¯")
		}
	}
}
