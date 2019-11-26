package main

import (
    "fmt"
)

func main() {

    firstRank := "39"
    secondRank := "614"

    if firstRank < secondRank {
        fmt.Println("First course is doing better than the second course")
    } else if firstRank > secondRank {
        fmt.Println("oh.... your first course is not doing well")
    } else {
        fmt.Println("Both course are same")
    }

}