package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}
type People struct {
	p []Person
}

func main() {
	p1 := Person{"navya", 20}
	p2 := Person{"durgesh", 21}
	p3 := Person{"rakshan", 10}
	p := []Person{p1, p2, p3}
	pe := People{p}
	//fmt.Println(p1, p2)
	sort.Slice(p, func(x, y int) bool {
		return p[x].age < p[y].age
	})
	fmt.Println(pe)
}
