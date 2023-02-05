package main

import (
)

type Word struct {
	name string
	edges []Word
}

func main() {
	_ = map[Word]bool{}
//	w1 := Word{name: "a", edges: []Word{}}
//	w2 := Word{name: "b", edges: []Word{w1}}
//	fmt.Printf("%p, %s\n", &w1, w1)
//	fmt.Printf("%p, %s\n", &w2.edges[0], w2.edges[0])
//	seen[w1] = true
//	fmt.Println(&w1 == &w2.edges[0])
	//fmt.Println(seen[w2.edges[0]])
}
