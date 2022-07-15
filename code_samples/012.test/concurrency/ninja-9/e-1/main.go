package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i :=0; i<4; i++{
		fmt.Println("I am first function",i)
		// wg.Wait()
		}
		wg.Done()
	}()
	fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())

	go func() {
		fmt.Println("I am second function")
		wg.Done()
	}()
	 fmt.Println("GoRoutine\t\t", runtime.NumGoroutine())
	 wg.Wait()
}
