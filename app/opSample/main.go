package main

import (
	"fmt"
)

type CC struct {
	Use     string
	use     string
	Comment string
	aa      int
}

func main() {
	fmt.Println("operation sample")
	s := "ssssssssss"
	a := &CC{Use: "yes", Comment: "for test"}
	b := &CC{Use: "yes", Comment: "for test"}
	fmt.Println("s's addr:", &s)
	fmt.Println("a's addr:", (&a), "b's addr:", b)
	fmt.Println("a==b:", &a == &b)
	fmt.Printf("a's addr:%x, b's addr:%x\n", a, b)
	fmt.Println("a==b:", a == b)
	fmt.Println("a==b:", *a == *b)
	a.use = "ok"
	fmt.Printf("a's addr:%v, b's addr:%v\n", a, b)
	fmt.Println("a==b:", a == b)
	b = a
	fmt.Printf("a's addr:%v, b's addr:%v\n", a, b)
	fmt.Println("a==b:", a == b)

	c := &CC{Use: "yes", Comment: "for test"}
	d := CC{Use: "yes", Comment: "for test"}
	fmt.Println("*c==d:", *c == d)
}
