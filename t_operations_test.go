package numstream

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(2, 3, 4, 5) != 14 {
		t.Errorf("Sum(2, 3, 4, 5) => %d, want 5", Sum(2, 3, 4, 5))
	}
}

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Errorf("Add(2, 3) => %d, want 5", Add(2, 3))
	}
}

func TestIncrementByX(t *testing.T) {
	i := IncrementByX(3)
	if i(5) != 8 {
		t.Errorf("IncrementByX(3)(5) => %d, want 8", IncrementByX(3)(5))
	}
}

func TestIncrement(t *testing.T) {
	a := Increment(4)
	if a != 5 {
		t.Errorf("Increment(4) => %d, want 5", a)
	}
}

var factorOutTests = []struct {
	factor, scalar, result int
}{
	{0, 7, 7},
	{2, 8, 1},
	{2, 9, 9},
	{5, 10, 2},
}

func TestFactorOut(t *testing.T) {
	for _, r := range factorOutTests {
		a := FactorOut(r.factor, r.scalar)
		if a != r.result {
			t.Errorf("FactorOut(%d, %d) => %d, want %d", r.factor, r.scalar, a, r.result)
		}
	}
}

func TestFactorOutX(t *testing.T) {
	for _, r := range factorOutTests {
		a := FactorOutX(r.factor)(r.scalar)
		if a != r.result {
			t.Errorf("FactorOutX(%d)(%d) => %d, want %d", r.factor, r.scalar, a, r.result)
		}
	}
}
