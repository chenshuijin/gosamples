package main

import (
	"fmt"
)

func QuickSort(list []int) {
	if list == nil {
		return
	}
	if len(list) <= 0 {
		return
	}

	mid, i := list[0], 1
	head, tail := 0, len(list)-1
	for head < tail {
		fmt.Println(list)
		if list[i] > mid {
			list[i], list[tail] = list[tail], list[i]
			tail--
		} else {
			list[i], list[head] = list[head], list[i]
			i++
		}
	}
	list[head] = mid
	QuickSort(list[:head])
	QuickSort(list[head+1:])
}
