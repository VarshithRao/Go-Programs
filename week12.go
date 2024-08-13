package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Germany", "G"))   // true
	fmt.Println(strings.ContainsAny("Germany", "g"))   // false
	fmt.Println(strings.Contains("Germany", "Ger"))    // true
	fmt.Println(strings.Contains("Germany", "ger"))    // false
	fmt.Println(strings.Contains("Germany", "er"))     // true
	fmt.Println(strings.Count("cheese", "e"))          // 3
	fmt.Println(strings.EqualFold("Cat", "cAt"))       // true
	fmt.Println(strings.EqualFold("India", "Indiana")) // false
}
