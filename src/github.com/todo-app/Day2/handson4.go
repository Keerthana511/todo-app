/*
create your own type.underlying type be an int
create identifier x with new type using var.
In func main,print the value,typeof x
assign 42 to x using"=" operator
print the value of x.
*/

package main

import (
	"fmt"
)

type flow int

var x flow

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
