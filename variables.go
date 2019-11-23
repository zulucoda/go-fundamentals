package main

import (
    "fmt"
    "reflect"
)

func main() {

    name := "Muzi"
//     course := "Golang"
    module := 1.5
    ptr := &module

    fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
    fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
    fmt.Println("Memory address of *module* variable is",
    ptr, "and the value of *module*", *ptr)
//     fmt.Println("Course is set to", course)
}