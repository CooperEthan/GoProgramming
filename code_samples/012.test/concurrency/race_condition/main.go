// go run -race main.go  --> to check race condition
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	// counter := 0
	var counter int64
	const gs = 100
	wg.Add(gs)

	// mutex --> to pretend race condition in the code.
	// var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func ()  {
			// v := counter
			//time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() 
			fmt.Println("Counter\t",atomic.LoadInt64(&counter))
			// v++
			// counter = v
			// mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())
	fmt.Println("count:",counter)
}