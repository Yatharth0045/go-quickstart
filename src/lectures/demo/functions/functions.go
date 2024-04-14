package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(a, b int) int {
	return a + b
}

func greet() {
	fmt.Println("Hello from Yatharth")
}

func main() {
	greet()
	dob := double(7)
	sum := add(double(10), 20)

	fmt.Println("Sum is", sum, "Double is", dob)
}
