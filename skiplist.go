package skiplist

import (
	"math"
	"time"
)

// Set updates the value that matches the key you already have
// When key is missing, occur error
func (s *SkipList) Set(key float64, value interface{}) *Element {
	update := s.getPrev(key)

	var element *Element

	if element = update[0].next[0]; element != nil && element.key == key {
		element.value = value
		return element
	}

	element = &Element{
		Node: Node{
			next: make([]*Element, s.maxHeight),
		},
		key:   key,
		value: value,
	}

	for i := range element.next {
		element.next[i] = update[i].next[i]
		update[i].next[i] = element
	}
	return element
}

// Delete removes Node that matches the key
func (s *SkipList) Delete(key float64) *Element {
	update := s.getPrev(key)

	if element := update[0].next[0]; element != nil && element.key == key {
		for key, value := range element.next {
			update[key].next[key] = value
		}
		return element
	}

	return nil
}

// Find get skiplist node when has same key
func (s *SkipList) Find(key float64) *Element {
	var prev *Node = &s.Node
	var next *Element

	for i := s.maxHeight - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && key > next.key {
			prev = &next.Node
			next = next.next[i]
		}
	}

	if next != nil && next.key == key {
		return next
	}

	return nil
}

// getPrev is the search previous node
func (s *SkipList) getPrev(key float64) []*Node {
	var prev *Node = &s.Node
	var next *Element

	var update = make([]*Node, s.maxHeight)

	for i := s.maxHeight - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && key > next.key {
			prev = &next.Node
			next = next.next[i]
		}
		update[i] = prev
	}
	return update
}

// generateRandomHeight generate level for new Node
func (s *SkipList) generateRandomHeight() int32 {
	s.rand.Lock()
	rnd := uint32(s.rand.src.Uint64())
	s.rand.Unlock()

	h := int32(1)
	for h < s.maxHeight && rnd <= s.probTable[h] {
		h++
	}

	return h
}

// probabilityTable generate table for generate random height
func probabilityTable(maxHeight int32, probability float64) (table []uint32) {
	p := float64(1.0)
	for i := int32(0); i < maxHeight; i++ {
		table = append(table, uint32(float64(math.MaxUint32)*p))
		p *= probability
	}
	return table
}

// NewWithDefault return a new empty SkipList
func NewWithDefault() *SkipList {
	skl := &SkipList{
		Node:        Node{next: make([]*Element, defaultMaxHeight)},
		maxHeight:   defaultMaxHeight,
		probability: defaultProbability,
		probTable:   probabilityTable(defaultMaxHeight, defaultProbability),
	}
	skl.rand.src.Seed(uint64(time.Now().UnixNano()))
	return skl
}

// New return a new empty SkipList
func New(maxHeight int32, probability float64) *SkipList {
	skl := &SkipList{
		Node:        Node{next: make([]*Element, maxHeight)},
		maxHeight:   maxHeight,
		probability: probability,
		probTable:   probabilityTable(maxHeight, probability),
	}
	skl.rand.src.Seed(uint64(time.Now().UnixNano()))
	return skl
}
