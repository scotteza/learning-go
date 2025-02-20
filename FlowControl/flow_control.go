package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number: ")
	text, _ := reader.ReadString('\n')
	stripped := strings.TrimSpace(text)
	fmt.Printf("You entered %v\n", stripped)
	someInt, err := strconv.Atoi(stripped)
	if err != nil {
		fmt.Println("Error converting string to integer")
		return
	}
	switch someInt % 10 {
	case 0:
		fmt.Println("The number is divisible by 10")
	default:
		fmt.Println("The number is not divisible by 10")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func infiniteLoop(sum int) {
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
