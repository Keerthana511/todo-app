/*for loop using for{} syntax
print out the years you been alive
*/
package main

import (
	"fmt"
)

func main() {
	bd := 1997
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
