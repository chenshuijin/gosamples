package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("go")
	fmt.Println("go")
loop:
	for {
		in := ""
		fmt.Scanln(&in)
		if in == "q" {
			break loop
		}
		fmt.Println("input:", in)
	}
}
