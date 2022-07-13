package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	xincrementer := 0
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := xincrementer
			runtime.Gosched()
			v++
			xincrementer = v
			fmt.Println(xincrementer)
			wg.Done()
		}()
		fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("end value:", xincrementer)
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())
	wg.Done()
}
