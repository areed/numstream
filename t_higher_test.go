package numstream

import (
	"testing"
)

func TestPartial2(t *testing.T) {
	add3 := Partial2(Add, 3)
	if add3(4) != 7 {
		t.Errorf("Partial2(Sum, 3)(4) => %d, want 7", Partial2(Add, 3)(4))
	}
}

func TestReduceUnary(t *testing.T) {
	o := ReduceUnary(func(n int) int { return n + 1 })(3, 5) //only the first argument is used
	if o != 4 {
		t.Fail()
	}
}

func TestReduceBinary(t *testing.T) {
	o := ReduceBinary(Add)(3, 5, 7) //only the first two arguments are used
	if o != 8 {
		t.Fail()
	}
}
