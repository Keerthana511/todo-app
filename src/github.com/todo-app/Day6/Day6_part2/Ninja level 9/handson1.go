/*In addition to main goroutine,launches two additional go routines
each additional goroutine prints something,
use wait grps
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("hello from thing 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("hello from thing 2")
		wg.Done()
	}()
	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("all about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

}
