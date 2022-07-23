package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Age":20},{"First":"Bond","Age":25}]`
	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Age)
	}

	fmt.Println("=========================")
	erro := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println("Error:", erro)
	}

	fmt.Println("=======================")
	// sort

}
