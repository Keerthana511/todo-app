/*
using prev code ,creating a variable y using var.
the variable should be of underlying type of your custom type ="x"
use the short declaration to assign that value to y.
printout the value,type of y.
*/
package main

import (
	"fmt"
)

type flow int

var x flow
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

}
