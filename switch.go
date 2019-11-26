package main

import (
    "fmt"
)

func main() {

    switch "docker" {
     case "linux":
        fmt.Println("Here are some recommended Linux courses....")
     case "docker":
        fmt.Println("Here are some recommended Docker courses....")
     default:
        fmt.Println("Sorry no match found....")
    }

}