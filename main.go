package main
import (
	"fmt"
	"time"
	"./packages"
	"./imports"
	"./functions"
	"./swap"
	"./named-results"
	"./variables"
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
	// exported names
	fmt.Println(imports.Math_pi())
	// functions
	fmt.Println(functions.Add(42, 13))
	// swap
	a, b := swap.Swap("hello", "world")
	fmt.Println(a, b)
	// named return values
	fmt.Println(named_results.Split(17))
	// variables
	variables.Variables()
	// variables with initializers
	variables.VariablesWithInitializers()
	// Short variable declarations
	variables.ShortVariableDeclarations()
}
