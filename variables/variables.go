package variables
import "fmt"

/*
	Variables
	@see https://tour.golang.org/basics/8
*/
var c, python, java bool

func Variables() {
	var i int
	fmt.Println(i, c, python, java)
}

/*
	Variables with initializers
	@see https://tour.golang.org/basics/9
*/
var i, j int = 1, 2
func VariablesWithInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

/*
	Short variable declarations
	@see https://tour.golang.org/basics/10
*/
func ShortVariableDeclarations() {
	k := 3
	c, python, java := true, false, "no!"
	
	fmt.Println(i, j, k, c, python, java)
}
