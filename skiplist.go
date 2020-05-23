package main

import (
	"fmt"
	"math"
)

const (
	defaultMaxHeight   int32   = 12
	defaultProbability float64 = 1 / math.E
)

// SkipList is a concurrent skiplist
type SkipList struct {
	maxHeight   int32
	probability float64
}

func genereateRandomHeight() {

}

// New return a new empty
func New() SkipList {
	return SkipList{
		maxHeight:   defaultMaxHeight,
		probability: defaultProbability,
	}
}

func main() {
	sl := New()
	fmt.Print(sl)
}
