package skiplist

import (
	"math"
	"time"
)

// generateRandomHeight generate level for new Node
func (s *SkipList) generateRandomHeight() uint32 {
	s.rand.Lock()
	rnd := uint32(s.rand.src.Uint64())
	s.rand.Unlock()

	h := uint32(1)
	for h < s.maxHeight && rnd <= s.probTable[h] {
		h++
	}

	return h
}

// probabilityTable generate table for generate random height
func probabilityTable(maxHeight uint32, probability float64) (table []uint32) {
	p := float64(1.0)
	for i := uint32(0); i < maxHeight; i++ {
		table = append(table, uint32(float64(math.MaxUint32)*p))
		p *= probability
	}
	return table
}

// NewWithDefault return a new empty SkipList
func NewWithDefault() *SkipList {
	skl := &SkipList{
		maxHeight:   defaultMaxHeight,
		probability: defaultProbability,
		probTable:   probabilityTable(defaultMaxHeight, defaultProbability),
	}
	skl.rand.src.Seed(uint64(time.Now().UnixNano()))
	return skl
}

// New return a new empty SkipList
func New(maxHeight uint32, probability float64) *SkipList {
	skl := &SkipList{
		maxHeight:   maxHeight,
		probability: probability,
		probTable:   probabilityTable(maxHeight, probability),
	}
	skl.rand.src.Seed(uint64(time.Now().UnixNano()))
	return skl
}
