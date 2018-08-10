package main

import (
	"fmt"
	"sort"
)

func main() {
	// go start a goroutine
	a := []int{1, 2, 3, 5, 6, 4, 2, 1, 2, 4, 5, 4, 9}
	sort.Ints(a)

	for i, v := range a {
		fmt.Println(i, v)
	}
}
