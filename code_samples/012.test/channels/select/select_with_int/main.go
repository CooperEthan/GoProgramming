package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send channel
func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// receive channel
func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}



// package main

// import "fmt"

// func main() {
// 	eve := make(chan int)
// 	odd := make(chan int)
// 	quit := make(chan int)

// 	// send
// 	go send(eve, odd, quit)

// 	// recieve
// 	recieve(eve, odd, quit)

// 	fmt.Println("about to exit")
// }

// func recieve(e, o, q <-chan int) {
// 	for {
// 		select {
// 		case v := <-e:
// 			fmt.Println("from the eve channle:", v)
// 		case v := <-o:
// 			fmt.Println("from the odd channle:", v)
// 		case v := <-q:
// 			fmt.Println("from the quit channle:", v)
// 			return
// 		}
// 	}
// }

// func send(e, o, q chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			e <- i
// 		} else {
// 			o <- i
// 		}
// 	}
// 	// close(e)
// 	// close(o)
// 	q <- 0
// }
