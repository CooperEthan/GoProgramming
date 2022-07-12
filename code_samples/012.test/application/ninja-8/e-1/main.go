package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type user struct {
	First string
	Age int
}

type First []user

func (a First) Len() int           { return len(a) }
func (a First) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a First) Less(i, j int) bool { return a[i].First < a[j].First }

type Age []user

func (a Age) Len() int           { return len(a) }
func (a Age) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Age) Less(i, j int) bool { return a[i].Age < a[j].Age }


func main() {

	u1 := user{
		First: "James",
		Age: 20,
	}
	u2 := user{
		First: "Bond",
		Age: 25,
	}
	users := []user{u1,u2}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error is: ",err)
	}
	fmt.Println(string(bs))

	//
	sort.Sort(First(users))
	for i,v := range users {
		fmt.Println("\nPERSON NUMBER",i)
		fmt.Println(v.First,v.Age)
	}

	fmt.Println("===============")

	sort.Sort(Age(users))
	for i,v := range users {
		fmt.Println("\nPERSON NUMBER",i)
		fmt.Println(v.First,v.Age)
	}

	
}