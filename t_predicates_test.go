package numstream

import (
	"testing"
)

var ints = []int{-99, -2, -1, 0, 1, 2, 99}

func TestTrue(t *testing.T) {
	for _, n := range ints {
		if !True(n) {
			t.Errorf("True(%d) => %v, want %v", n, false, true)
		}
	}
}

func TestFalse(t *testing.T) {
	for _, n := range ints {
		if False(n) {
			t.Errorf("False(%d) => %v, want %v", n, true, false)
		}
	}
}

func TestDisjoin(t *testing.T) {
	var tests = []struct {
		in  []Predicate
		out bool
	}{
		{[]Predicate{}, false},
		{[]Predicate{True}, true},
		{[]Predicate{False}, false},
		{[]Predicate{True, False}, true},
		{[]Predicate{False, True}, true},
		{[]Predicate{False, False, False}, false},
		{[]Predicate{False, True, False}, true},
	}
	for i, r := range tests {
		p := Disjoin(r.in...)
		for _, n := range ints {
			b := p(n)
			if b != r.out {
				t.Errorf("Disjoin([row %d Predicates]...) => %v, want %v", i, b, r.out)
			}
		}
	}
}

func TestConjoin(t *testing.T) {
	var tests = []struct {
		in  []Predicate
		out bool
	}{
		{[]Predicate{}, true},
		{[]Predicate{True}, true},
		{[]Predicate{False}, false},
		{[]Predicate{True, False}, false},
		{[]Predicate{False, True}, false},
		{[]Predicate{True, True}, true},
		{[]Predicate{True, False, True}, false},
		{[]Predicate{True, True, True}, true},
		{[]Predicate{False, False, False}, false},
	}
	for i, r := range tests {
		p := Conjoin(r.in...)
		for _, n := range ints {
			b := p(n)
			if b != r.out {
				t.Errorf("Conjoin([row %d Predicates]...) => %v, want %v", i, b, r.out)
			}
		}
	}
}

func TestDone(t *testing.T) {
	ch := make(chan bool)
	isDone := Done(ch)
	for _, n := range ints {
		if isDone(n) {
			t.Errorf("isDone(%d) before signal sent on done channel => false, want true", n)
		}
	}
	ch <- true
	for _, n := range ints {
		if !isDone(n) {
			t.Errorf("isDone(%d) after signal sent on done channel => true, want false", n)
		}
	}
}

func TestIsLessThan(t *testing.T) {
	b := IsLessThan(3, 4)
	if b {
		t.Errorf("IsLessThan(3, 4) => true, want false")
	}
}

func TestIsLessThanX(t *testing.T) {
	isLessThan10 := IsLessThanX(10)
	if isLessThan10(11) {
		t.Errorf("IsLessThan(10)(11) => true, want false")
	}
}

func TestIsMultipleOf(t *testing.T) {
	b := IsMultipleOf(3, 6)
	if !b {
		t.Error("6 is a multiple of 3")
	}
}

func TestIsMultipleOfX(t *testing.T) {
	var tests = []struct {
		multiple int
		value    int
		out      bool
	}{
		{2, 2, true},
		{2, 3, false},
		{2, 4, true},
		{3, 1, false},
		{3, 2, false},
		{3, 3, true},
		{3, 4, false},
		{3, 6, true},
		{3, 9, true},
		{10, 200, true},
		{4, 1024, true},
	}
	for _, r := range tests {
		p := IsMultipleOfX(r.multiple)
		b := p(r.value)
		if b != r.out {
			t.Errorf("IsMultipleOf(%d)(%d) => %v, want %v", r.multiple, r.value, b, r.out)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		in  int
		out bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
		{7, true},
		{8, true},
		{9, true},
		{10, false},
		{11, true},
		{12, false},
		{99, true},
		{100, false},
		{101, true},
		{9889, true},
		{9989, false},
		{6543456, true},
		{10999901, true},
		{76535677, false},
	}
	for _, r := range tests {
		b := IsPalindrome(r.in)
		if b != r.out {
			t.Errorf("IsPalindrome(%d) => %v, want %v", r.in, b, r.out)
		}
	}
}
