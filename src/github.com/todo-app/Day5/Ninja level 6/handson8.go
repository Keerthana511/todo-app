/*fn that returns a func,
assign returned func to a variable,
call that func
*/
package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())
}
func foo() func() int {
	return func() int {
		return 42
	}
}
