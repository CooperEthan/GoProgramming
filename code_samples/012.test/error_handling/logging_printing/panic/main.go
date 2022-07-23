package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("no-file.txt")
	if err != nil {
		log.Panic(err)
		// panic(err)
	}
	defer f.Close()
}
