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
	// var xincrementer int64
	var mu sync.Mutex
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := xincrementer
			// runtime.Gosched()
			v++
			xincrementer = v

			fmt.Println(xincrementer)
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("end value:", xincrementer)
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())
	// wg.Done()  --> dont put done() here
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	incrementer := 0
// 	gs := 100
// 	wg.Add(gs)
// 	var m sync.Mutex

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			m.Lock()
// 			v := incrementer
// 			v++
// 			incrementer = v
// 			fmt.Println(incrementer)
// 			m.Unlock()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("end value:", incrementer)
// }
