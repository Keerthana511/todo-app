/*directional channel(send only)
 */
package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println("------\n")
	fmt.Printf("cs\t%T\n", cs)

}
