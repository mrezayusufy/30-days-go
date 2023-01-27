package main

import "fmt"

func main() {
	var myName = "Reza Yusufy"
	fmt.Println("My name is ", myName)

	var name string = "Reyu"
	fmt.Println("Name : ", name)

	userName := "Administrator"
	fmt.Println("User Name : ", userName)

	var sum int
	fmt.Println("The sum is :", sum)

	part1, other := 1, 5
	fmt.Println("Part 1: ", part1, ", Other: ", other)

	part2, other := 2, 0
	fmt.Println("part 2: ", part2, ", Other: ", other)

	sum = part1 + part2
	fmt.Println("The sum is :", sum)

	var (
		lessonName = "Lesson: Variables"
		lessonType = "Lesson Type: Demo"
	)

	fmt.Println("Lesson Name: ", lessonName)
	fmt.Println("Lesson Type: ", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

}
