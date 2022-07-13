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
	var xincrementer int64
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt64(&xincrementer, 1)
			runtime.Gosched()
			fmt.Println("Atomic\t", atomic.LoadInt64(&xincrementer))

			wg.Done()
		}()
		fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("end value:", xincrementer)
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())
	// wg.Done()  --> dont put done() here
}
