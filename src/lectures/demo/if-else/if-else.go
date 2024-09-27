package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 scored higher")
	} else if quiz1 < quiz2 {
		fmt.Print("Quiz2 scored higher")
	} else {
		fmt.Println("Scores are equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Grades accepted")
	} else {
		fmt.Println("Grades rejected")
	}
}
