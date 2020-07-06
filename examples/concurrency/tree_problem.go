package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	//fmt.Println(t)
	ch <- t.Value
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree, ch1, ch2 chan int) bool {
	var v, v1, v2 []int
	m := map[*tree.Tree]chan int{t1:ch1, t2:ch2}
	
	for t, ch := range m {
		go Walk(t,ch)
		//t := false
		for {
			if len(v) == 10 {
				sort.Ints(v)
				if t == t1 {
					v1 = v
				} else { 
					v2 = v 
				}
				fmt.Println("Values: ", v)
				v = []int{}
				break
			}
			v = append(v, <- ch)
			//fmt.Println(v)
		}
	}
	for i, v := range v1 {
		for j, w := range v2{
			if i ==j && v != w {
				return false
			}
		}
	}
	return true
}


func main() {
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	//go Walk(tree.New(1), ch)
	fmt.Println(Same(tree.New(1), tree.New(1), ch1, ch2))
	fmt.Println(Same(tree.New(1), tree.New(2), ch1, ch2))
	
}
