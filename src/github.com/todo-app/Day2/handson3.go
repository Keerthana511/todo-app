/*
At the package level scope,assign values.
use fmt.SPrint to all of the values to one single string.
*/

package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
