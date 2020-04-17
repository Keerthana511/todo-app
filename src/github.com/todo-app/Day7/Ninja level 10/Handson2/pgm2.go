/*directional channel(recieve only)
*/
package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)
	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)
	
}
