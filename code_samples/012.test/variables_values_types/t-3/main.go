package main

import "fmt"


var y = 42
var z = "s"

func main(){

fmt.Println(y)
fmt.Printf("%T\n",y)
fmt.Println(z)

z= "new s"
fmt.Println(z)

}