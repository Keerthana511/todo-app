/*use mutex to fix prev pgm
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
