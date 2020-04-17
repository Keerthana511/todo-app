/*pgm that launches 10 go routines
each go routines adds 10 nos to a channel
pull the no's off the channel and print them
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
    for j:=0;j<10;j++{
        go func(){
            for i:=0;i<10;i++{
                c<-i
            }
        }()
    }
    for k:=0;k<100;k++{
        fmt.Println(k, <-c)
    }
    fmt.Println("all about to exit")
    }