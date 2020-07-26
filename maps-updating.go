package main

import (
	"fmt"
)

// Maps Similar to arrays and slices but:
// 1) Maps are unordered
// 2) Maps key:value pair (JavaScript object)
// 3) Maps Dynamically resizable
// 4) Maps are referecnes
// 5) Maps are NOT safe for concurrency
// 6) specify size for large maps

func main() {

	// declare map[keType]valueType
	myAlfaRomeos := make(map[string]int)

	myAlfaRomeos["4c"] = 2
	myAlfaRomeos["giulia"] = 3
	myAlfaRomeos["stelvio"] = 4
	myAlfaRomeos["gtam"] = 1

	fmt.Println(myAlfaRomeos)

	recentTrackDays := map[string]int{
		"4c":   0,
		"gtam": 5,
	}

	fmt.Println(recentTrackDays)

	fmt.Println(myAlfaRomeos["gtam"])

	myAlfaRomeos["4c"] = 10
	fmt.Println(myAlfaRomeos)

}
