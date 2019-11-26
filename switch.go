package main

import (
    "fmt"
    "math/rand"
    "time"
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


    main2()

}

func main2() {

    switch tmpNum := random(); tmpNum {
     case 0,2,4,6,8:
        fmt.Println("We got an even number", tmpNum)
     case 1,3,5,7,9:
        fmt.Println("We got an odd number", tmpNum)
    }

}

func random() int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(10)
}