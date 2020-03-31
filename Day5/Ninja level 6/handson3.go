/*using defer keyword
*/
package main

import (
	"fmt"
)

func main(){
    defer foo()
    fmt.Println("Hello ,Keerthana")
}
func foo(){
    defer func(){
        fmt.Println("foo defer ran")
    }()
    fmt.Println("foo ran")
}