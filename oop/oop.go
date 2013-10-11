package oop

import (
	"fmt"
)

const (
	TEST_STRING = "123"
)

func Init() {
	fmt.Println(TEST_STRING)
}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Xi []int
type Si []string

func (p Xi) Len() int {
	return len(p)
}

func (p Xi) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p Xi) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (s Si) Len() int {
	return len(s)
}

func (s Si) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Si) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}
