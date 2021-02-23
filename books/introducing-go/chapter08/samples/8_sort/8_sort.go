package main

import (
	"fmt"
	"sort"
)

func main() {
	kids := []Person{
		{"Katrin", 11},
		{"Stephani", 8},
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}

type Person struct {
	Name string
	Age int
}

type ByName []Person

// The implementation of sort.Interface
func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person

func (ba ByAge) Len() int {
	return len(ba)
}

func (ba ByAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}



