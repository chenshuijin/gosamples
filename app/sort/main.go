package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("yes")
	a := []int{90, 43, 21, 43, 21, 436, 23, 673, 12, 343}
	sort.Ints(a)
	fmt.Println(a)
	for i := 10; i > 0; i >>= 1 {
		fmt.Println("i:", i)
	}

}

func divid_sort(s []int) []int {
	sLen := len(s)

	switch {
	case sLen <= 1:
		return s
	case sLen == 2:
		ret := []int{}
		if s[0] > s[1] {
			ret[0], ret[1] = s[1], s[0]
		} else {
			ret[0], ret[1] = s[0], s[1]
		}
		return ret
	default:
		q := len(s) / 2
		l, r := []int{}, []int{}
		copy(l, s[:q])
		copy(r, s[q:])
		divid_sort(l)
	}
	return nil
}

type Int []int

func (self Int) Less(i, j int) bool {
	return self[i] > self[j]
}
