 package main

 import (
     "fmt"
 )

 func main() {

    coursesInProgress := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}
    coursesInCompleted := []string{"Docker Deep Dive", "Go Fundamentals"}

    for _, i := range coursesInProgress {
        fmt.Println(i)
        for _, j := range coursesInCompleted {
            if i == j {
                fmt.Println(i, "is in both lists")
            }
        }
    }
 }