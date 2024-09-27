package main

import "fmt"

func main() {
	myName := "Yatharth"
	// println automtically adds space when there is a comma
	fmt.Println("My name is", myName, myName)

	var name string = "Yash"
	fmt.Println("name =", name)

	var sum int
	fmt.Println("Sum is", sum)

	part1, part2 := 1, 5
	fmt.Println("part1 is", part1, "Other is", part2)

	x, y := 3, 5
	z, y := 7, 9
	fmt.Println("x is", x, "y is", y, "z is", z)

	var (
		lessonName = "Maths"
		lessonType = "Nice"
	)

	fmt.Println("lessonName is", lessonName, "lessonType is", lessonType)

	w1, w2, _ := "x", "y", "z"
	fmt.Println(w1, w2)
}
