package main
import (
	"fmt"
	"time"
	"math"
	"./packages"
	"./imports"
	"./functions"
)

func main() {
	// print
	fmt.Println("Hello, World!")
	// time
	fmt.Println("The time is", time.Now())
	// packages
	fmt.Println("My favorite number is", packages.Pack())
	// imports
	fmt.Println("Now you have problems.\n", imports.Import())
	/*
		exported names
		@see https://tour.golang.org/basics/3
	*/
	fmt.Println(math.Pi)
	// functions
	fmt.Println(functions.Add(42, 13))
}
