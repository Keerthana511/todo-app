/*print remainders b/w 10 &100,when it is divides by 4
*/
package main
import (
    "fmt"

)
func main(){
    for i:=10;i<=100;i++{
        fmt.Printf("%v\n",i%4)
    }
}