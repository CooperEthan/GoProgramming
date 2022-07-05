package main

import "fmt"


func main() {

	m := map[string][]string {
		"cooper_ethan":{"soccer","shooting"},
		"akk_hid":{"basketball","footbaall"},
	}

	fmt.Println("===================")

	m["b_a"] = []string{"film","movie"}
	fmt.Println(m)

	for k,v := range m {
		fmt.Println(k)
		for j,t := range v {
			fmt.Println(j,t)
		}
	}

	fmt.Println("===================")

	delete(m, "b_a")

	for k,v := range m {
		fmt.Println(k)
		for j,t := range v {
			fmt.Println(j,t)
		}
	}

}