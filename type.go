package skiplist

import (
	"math"
	"sync"

	"golang.org/x/exp/rand"
)

const (
	defaultMaxHeight   int32   = 12
	defaultProbability float64 = 1 / math.E
)

// Node represents a Node in skiplist
type Node struct {
	next []*Element
}

// Element is struct consisting of key & value
type Element struct {
	Node
	key   float64
	value interface{}
}

// Key is identifier
func (e *Element) Key() float64 {
	return e.key
}

// SkipList is a concurrent skiplist
type SkipList struct {
	Node
	maxHeight   int32
	probability float64
	probTable   []uint32

	rand struct {
		sync.Mutex
		src rand.PCGSource
	}
}
