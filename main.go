package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	age  int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {
	kids := []Person{
		{"duc", 9},
		{"anh", 12},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
