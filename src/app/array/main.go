package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5, 6}
	b := a
	b[2] = 77
	fmt.Println(a, b)

	a1 := [5]int{2, 3, 4, 5, 6}
	b1 := a1
	b1[2] = 77
	fmt.Println(a1, b1)
}
