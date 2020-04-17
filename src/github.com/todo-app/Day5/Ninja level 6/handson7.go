/*assign a func to a variable,call that func
 */
package main

import (
	"fmt"
)

func main() {
	f := func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}
	f()
}
