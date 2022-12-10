package main

import (
	"fmt"

	"github.com/marcelcorso/ll"
)

func main() {
	l := ll.LL{}
	l.Insert(4)
	l.Insert(8)
	l.Insert(2)
	l.Insert(9)
	l.Insert(888)
	l.Insert(-1)
	l.Insert(54)

	l.Each(func(val int) {
		fmt.Print(val)
	})
	fmt.Println("")
}
