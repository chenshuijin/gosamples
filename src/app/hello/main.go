package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	SampleArgs1([]string{"h", "d", "daf", "sdf", "2", "afd", "defg"})
	SampleArgs2("h", "d", "daf", "sdf", "2", "afd", "defg")
}

func SampleArgs1(args []string) {
	fmt.Println("args count:", len(args))
	fmt.Println("args:", args)
	for arg := range args {
		fmt.Println(arg)
	}
}

func SampleArgs2(args ...string) {
	fmt.Println("args count:", len(args))
	fmt.Println("args:", args)
	for arg := range args {
		fmt.Println(arg)
	}
}
