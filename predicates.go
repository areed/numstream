package numstream

import (
	"fmt"
)

//a Predicate takes an int and returns a boolean.
type Predicate func(int) bool
type Predicate2 func(int, int) bool

//True always returns true.
func True(n int) bool {
	return true
}

//False always returns false.
func False(n int) bool {
	return false
}

//IsLess returns true iff the second argument is less than the first.
func IsLessThan(n, m int) bool {
	return m < n
}

//IsMultipleOf returns true iff the second argument is a multiple of the first.
func IsMultipleOf(m, n int) bool {
	return n%m == 0
}

//IsPalindrome returns true iff the int argument can be reversed without changing its value.
func IsPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	l := len(s)
	for i, c := range s {
		if c != rune(s[l-i-1]) {
			return false
		}
	}
	return true
}

/* Functions returning Predicates */

//Partial2Pred returns a Predicate from a Predicate2 with the first argument preapplied
func Partial2Pred(p Predicate2, m int) Predicate {
	return func(n int) bool {
		return p(m, n)
	}
}

//Done returns a Predicate that evaluates to true iff a signal has been received on the done channel.
func Done(done <-chan bool) Predicate {
	var d bool
	go func() {
		d = <-done
	}()
	return func(n int) bool {
		return d
	}
}

//Conjoin returns a Predicate that evalutes to true iff every Predicate argument evaluates to true
//for the applied int.
func Conjoin(preds ...Predicate) Predicate {
	return func(n int) bool {
		for _, p := range preds {
			if !p(n) {
				return false
			}
		}
		return true
	}
}

//Disjoin returns a Predicate that evaluates to true iff any Predicate argument evaluates to true
//for the applied int
func Disjoin(preds ...Predicate) Predicate {
	return func(n int) bool {
		for _, p := range preds {
			if p(n) {
				return true
			}
		}
		return false
	}
}

//IsLessThanX returns a Predicate that evalutes to true iff the applied int argument is less than the bound int argument.
func IsLessThanX(x int) Predicate {
	return Partial2Pred(IsLessThan, x)
}

//IsMultipeOf returns a Predicate that evaluates to true iff the applied int is a multiple of the bound int argument.
func IsMultipleOfX(x int) Predicate {
	return Partial2Pred(IsMultipleOf, x)
}
