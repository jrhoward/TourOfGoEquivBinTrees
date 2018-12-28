package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()
	var slice1, slice2 []int
	for {
		t1val, ch1ok := <-ch1
		if ch1ok {
			slice1 = append(slice1, t1val)
		}
		t2val, ch2ok := <-ch2
		if ch2ok {
			slice2 = append(slice2, t2val)
		}
		if !ch1ok && !ch2ok {
			break
		}
	}
	return Equal(slice1, slice2)
}

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		Walk(tree.New(1), ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Trees are same = ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Trees are same = ", Same(tree.New(1), tree.New(2)))
}
