package test

import (
	"fmt"
	"os"
)

func main() {
	// This is a single-line comment
	fmt.Println("Hello, World!")

	/*
		This is a multi-line comment
		It spans multiple lines
	*/

	var a int = 10
	var b int = 20
	var c int = a + b

	if c > 25 {
		fmt.Println("c is greater than 25")
	} else {
		fmt.Println("c is 25 or less")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	switch c {
	case 10:
		fmt.Println("c is 10")
	case 20:
		fmt.Println("c is 20")
	default:
		fmt.Println("c is neither 10 nor 20")
	}

	// Function call
	result := add(a, b)
	fmt.Println("Result of add function:", result)

	// Slice example
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println("Number:", num)
	}

	// Map example
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Error handling
	file, err := os.Open("nonexistentfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	} else {
		fmt.Println("File opened successfully")
		file.Close()
	}
}

func add(x int, y int) int {
	return x + y
}
