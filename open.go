package main

import (
    "fmt"
    "os"
)

func main() {

    _, err := os.Open("some file")

    if err != nil {
        fmt.Println("Error returned was:", err)
    }

}