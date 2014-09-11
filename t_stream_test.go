package numstream

import (
	"testing"
)

func TestMerge(t *testing.T) {
	a := ToChannel([]int{0, 2, 4, 6, 8})
	b := ToChannel([]int{1, 3, 5, 7, 9})
	c := Merge(a, b)
	m := make([]bool, 10)
	for n := range c {
		m[n] = true
	}
	for i, v := range m {
		if !v {
			t.Errorf("Expected merged stream to contain %d", i)
		}
	}
}
