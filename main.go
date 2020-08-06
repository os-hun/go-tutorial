package main
import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

/*
	functions
	@see https://tour.golang.org/basics/4
*/
func add(x int, y int) int {
	return x + y
}

func main() {
	// print
	fmt.Println("Hello, World!")
	// time
	fmt.Println("The time is", time.Now())
	/*
		packages
		@see https://tour.golang.org/basics/1
	*/
	fmt.Println("My favorite number is", rand.Intn(20))
	/*
		imports
		@see https://tour.golang.org/basics/2
	*/
	fmt.Println("Now you have problems.\n", math.Sqrt(7))
	/*
		exported names
		@see https://tour.golang.org/basics/3
	*/
	fmt.Println(math.Pi)
	// functions
	fmt.Println(add(42, 13))
}
