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

func TestInsertAndFind(t *testing.T) {
	s := NewWithDefault()
	testDataInsert := []float64{3, 4}
	testDataNotInsert := []float64{1, 5}

	for _, key := range testDataInsert {
		s.Insert(key, 0)
		fmt.Printf("%f key insert\n", key)
	}

	for _, key := range testDataInsert {
		if s.Find(key) == nil {
			t.Errorf("Can't find added key")
		} else {
			fmt.Printf("%f key exists\n", key)
		}
	}

	for _, key := range testDataNotInsert {
		if s.Find(key) != nil {
			t.Errorf("Can't find not added key")
		} else {
			fmt.Printf("%f key doesn't exists\n", key)
		}
	}
}

func TestInsertAndDeleteAndFind(t *testing.T) {
	s := NewWithDefault()
	testDataInsert := []float64{3, 4, 5}
	testDataDelete := []float64{4}
	testDataExists := []float64{3, 5}
	testDataNotExists := []float64{4}

	for _, key := range testDataInsert {
		s.Insert(key, 0)
		fmt.Printf("%f key insert\n", key)
	}

	for _, key := range testDataDelete {
		if s.Delete(key) == nil {
			t.Errorf("Can't delete key")
		}
	}

	for _, key := range testDataExists {
		if s.Find(key) == nil {
			t.Errorf("Can't find added key")
		} else {
			fmt.Printf("%f key exists\n", key)
		}
	}

	for _, key := range testDataNotExists {
		if s.Find(key) != nil {
			t.Errorf("Should not find deleted key")
		} else {
			fmt.Printf("%f key deleted\n", key)
		}
	}
}
