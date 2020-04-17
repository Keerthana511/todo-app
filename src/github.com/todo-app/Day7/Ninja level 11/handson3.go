/*create a struct "customErr" which implements builtin interface
create a value of type that type and pass it to foo
*/
package main

import (
	"fmt"
)
type customErr struct{
    info string
}
func (ce customErr) Error() string{
    return fmt.Sprintf("here is the error:%v", ce.info)
}
func main(){
    c1:=customErr{
        info:"need more coffee",
    }
    foo(c1)
}
func foo(e error){
    fmt.Println("foo ran -", e, "\n")
}