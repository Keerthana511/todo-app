/*create an identifier that returns an identifier an int
create an identifier that returns both int and string
call both fn's,print out the results.
*/
package main

import (
	"fmt"
)

func main() {
	n := foo()
	x, s := bar()
	fmt.Println(n, x, s)
}
func foo() int {
	return 42
}
func bar() (int, string) {
	return 1984, "big bro is watching"
}
