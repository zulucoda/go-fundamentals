package main

import (
    "fmt"
    "os"
)

func main() {

    name := os.Getenv("USERNAME")
    course := "Golang"

    fmt.Println("\nHi", name, "you're currently watching", course)

    changeCourse(&course)

    fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {
    *course = "First Look: Native Docker Clustering"

    fmt.Println("\nTrying to changer your course to", *course)

    return *course
}