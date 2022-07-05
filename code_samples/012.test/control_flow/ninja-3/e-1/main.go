package main

import "fmt"


func main() {
	// for i := 0; i <= 10000; i++ {
	// 	fmt.Println(i)
	// }
//=================
	// for i := 65; i <= 90; i++ {
	// 	fmt.Println("\n",i)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("%#U\n", i)
	// 	}
	// }
	//=====================
	// bd := 1981
	// for {	
	// 	if bd <= 2022 {
	// 		fmt.Println(bd)
	// 		bd++
	// 	}else {
	// 		break
	// 	}

	// }
	// for bd <=2022 {
	// 	fmt.Println(bd)
	// 	bd++
	// }

		//===================================

	// for i := 10; i <= 100; i++ {
	// 	// if i % 4 == 0 {
	// 	// 	fmt.Println(i)
	// 	// }
	// 	fmt.Println(i%4)
	// }
			//===================================
	
		i := 4

		if i >=4 {
			fmt.Println("it is >=4")
		}else if i<4 {
			fmt.Println("it is less than 4")
		}else{
			fmt.Println("nothing")
		}

					//===================================
		
		switch {
		case false:
			fmt.Println("no print")
		case true:
			fmt.Println("should print")
		}
						//===================================
		n := "favSport"
		switch n {
		case "ethan":
			fmt.Println("ethan")
		case "favSport":
			fmt.Println("yes it is!!!")
		default:
			fmt.Println("default case")
		}
							//===================================
		fmt.Println(true || false)
		fmt.Println(true && false)
		fmt.Println(false && false)
		fmt.Println(false || false)
}