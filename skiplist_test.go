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
