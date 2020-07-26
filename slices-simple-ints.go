package main

import (
	"fmt"
)

func main() {

	// Deckares a slice called myCourses
	myCourses := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(myCourses[4])

	myCourses[1] = 0
	fmt.Println(myCourses)

	fmt.Printf("Length is %d.\nCapacity is: %d", len(myCourses), cap(myCourses))
}
