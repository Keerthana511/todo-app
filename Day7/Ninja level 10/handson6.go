/*pgm that puts 100 nos to a channel
pull the nos off the channel and print them
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("all about to exit")
}
