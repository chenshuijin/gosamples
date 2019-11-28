package main

import (
	"flag"
	"fmt"
)

var p = flag.String("p", "a flag sample", "only a sample")

func main() {
	fmt.Println("flag sample")
	flag.Parse()
	fmt.Println("flag is p is:", *p)
}
