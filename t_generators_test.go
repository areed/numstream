package numstream

import (
	"testing"
)

func TestToSlice(t *testing.T) {
	ch := ToChannel([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	slice := ToSlice(ch)
	for i, n := range slice {
		if i != n {
			t.Fail()
		}
	}
}

func TestGenerate(t *testing.T) {
	ch := Generate(Increment, IsLessThanX(10), 0)
	slice := ToSlice(ch)
	for i, n := range slice {
		if i != n {
			t.Fail()
		}
	}
}
