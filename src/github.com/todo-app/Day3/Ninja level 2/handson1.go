/*pgm that prints a no in dec,binary,hex
 */
package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x", x)

}
