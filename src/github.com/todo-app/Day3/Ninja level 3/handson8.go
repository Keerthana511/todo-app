/*pgm  that uses switch statement but no switches specified
*/
package main
import (
    "fmt"

)
func main(){
    switch{
        case false:
            fmt.Println("should not print")
        case true:
            fmt.Println("should print")
    }
}