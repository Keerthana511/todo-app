/*for loop using for condition{} syntax
print out the years you been alive
*/
package main

import (
	"fmt"
)

func main() {
	bd := 1997
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}
}
