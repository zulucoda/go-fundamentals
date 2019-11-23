package main

import (
    "fmt"
    "reflect"
)

var (
    name, course, module = "Muzi", "Golang", 1.5
)

func main() {
    fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
    fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
    fmt.Println("Course is set to", course)
}