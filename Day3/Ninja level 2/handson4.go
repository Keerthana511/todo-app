/*assigns int to a variable
prints that int into dec,binary,hex
shifts the bit of int over 1 position to left and assign that to a variable
*/
package main

import (
	"fmt"
)

func main() {
	a := 40
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)

}
