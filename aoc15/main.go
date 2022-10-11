package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	numArgs := len(os.Args)
	if numArgs < 2 {
		log.Fatal("Not enough args.")
	}

	filename := os.Args[1]

	fmt.Println("File:", filename)
}
