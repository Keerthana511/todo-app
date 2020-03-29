/*pgm that uses swtch statement withe switch  expr,
specified as a variable type string with identifier "favSport"
*/
package main

import (
	"fmt"
)

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to mountain")
	case "surfing":
		fmt.Println("go to surfy")
	case "swimming":
		fmt.Println("go to pool")

	}
}
