package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle 
	fourwheel bool
}
type sedan struct{
	vehicle
	luxury bool
}


func main() {

	t := truck {
		vehicle: vehicle{
			doors : 5,
			color: "blue",
		},
		fourwheel: true,
	}
	// fmt.Println(t)

	s := sedan {
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(s)

	fmt.Println(t.fourwheel)
	fmt.Println(t.doors)
	fmt.Println(t.color)

}