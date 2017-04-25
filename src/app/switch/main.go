package main

import (
	"fmt"
	"os"
	"runtime"
)

func init() {

}
func main() {
	fmt.Println("args:", os.Args)
	fmt.Println("gomaxprocs:", runtime.NumCPU())
	switch {
	case true:
		fmt.Println("case1")
		fallthrough
	case true:
		fmt.Println("case2")
	}
}
