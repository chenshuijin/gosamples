package main

import (
	"bcd"
	"config"
	"fmt"
)

func main() {
	fmt.Println("flag sample")
	config.Parse("config parse")
	fmt.Println(bcd.Get())
}
