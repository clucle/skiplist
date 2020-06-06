package skiplist

import (
	"fmt"
	"testing"
)

func TestGenerateLevel(t *testing.T) {
	s := NewWithDefault()
	var levelDistribution [defaultMaxHeight + 1]int32
	for i := 0; i < 10000000; i++ {
		levelDistribution[s.generateRandomHeight()]++
	}
	fmt.Println(levelDistribution)
}

func TestSetAndFind(t *testing.T) {
	s := NewWithDefault()
	testDataSet := []float64{3, 4}
	testDataNotSet := []float64{1, 5}

	for _, key := range testDataSet {
		s.Set(key, 0)
		fmt.Printf("%f key set\n", key)
	}

	for _, key := range testDataSet {
		if s.Find(key) == nil {
			t.Errorf("Can't find added key")
		} else {
			fmt.Printf("%f key exists\n", key)
		}
	}

	for _, key := range testDataNotSet {
		if s.Find(key) != nil {
			t.Errorf("Can't find not added key")
		} else {
			fmt.Printf("%f key doesn't exists\n", key)
		}
	}
}
