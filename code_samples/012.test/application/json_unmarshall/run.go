package main

import (
	"encoding/json"
	"fmt"
)


type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	// fmt.Println(bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("Error: ",err)
	}
	for i,v := range people {
		fmt.Println("\nPERSON NUMBER",i)
		fmt.Println(v.First,v.Last,v.Age)
	}


}