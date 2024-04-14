//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Welcome", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func greetReturn(name string) string {
	return "Welcome " + name
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func sumOfThree(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func random() int {
	return 5
}

// * Write a function that returns any two numbers
func returnTwo() (int, int) {
	return 5, 8
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

func main() {
	greet("Yatharth")
	fmt.Println("Hey There,", greetReturn("Yatharth"))
	fmt.Println("Sum of digits is:", sumOfThree(10, 12, 24))
	a := random()
	fmt.Println("Returning number", a)
	b, c := returnTwo()
	fmt.Println("Returning two numbers:", b, c)
	fmt.Println("Sum of digits is:", sumOfThree(random(), b, c))
}
