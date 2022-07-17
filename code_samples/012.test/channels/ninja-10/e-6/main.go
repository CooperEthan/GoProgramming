package main

import "fmt"

func main() {

	n := make(chan int)
	m := gen(n)

	recieve(m, n)

	fmt.Println("about to exit")

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func recieve(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}

}
