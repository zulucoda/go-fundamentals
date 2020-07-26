package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// this was gives us a pointer
	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	// declare using composite literal way
	DockerDeepDive := courseMeta{
		Author: "Some Author Name 1",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println("Docker deep dive author is:", DockerDeepDive.Author)
}
