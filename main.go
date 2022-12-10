package main

import "fmt"

type LL struct {
	Head *llnode
}

type llnode struct {
	Val  int
	Next *llnode
}

func (ll *LL) Insert(val int) {
	n := &llnode{
		Val: val,
	}
	p := &ll.Head
	for {
		if *p == nil {
			*p = n
			return
		}

		if (**p).Val < val {
			p = &(**p).Next
		} else {
			n.Next = *p
			*p = n
			return
		}
	}
}

func (ll *LL) Each(f func(val int)) {
	for p := &ll.Head; *p != nil; p = &(*p).Next {
		f((**p).Val)
	}
}

func main() {
	ll := LL{}
	ll.Insert(4)
	ll.Insert(8)
	ll.Insert(2)
	ll.Insert(9)
	ll.Insert(888)
	ll.Insert(-1)
	ll.Insert(54)

	ll.Each(func(val int) {
		fmt.Print(val)
	})
	fmt.Println("")
}
