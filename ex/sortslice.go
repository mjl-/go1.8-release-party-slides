package main

import (
	"sort"
	"fmt"
)

type E struct {
	Value int
}

func main() {
	a := []E{
		{2},
		{3},
		{1},
		{0},
		{10},
		{5},
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Value < a[j].Value
	})
	fmt.Printf("%v\n", a)
}
