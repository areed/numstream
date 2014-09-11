package numstream

type Unary func(int) int
type Binary func(int, int) int
type Reduce func(ns ...int) int

//Partial2 applies a single int argument to a Binary and returns a Unary
func Partial2(f Binary, n int) Unary {
	return func(m int) int {
		return f(n, m)
	}
}

func ReduceUnary(f Unary) Reduce {
	return func(ns ...int) int {
		return f(ns[0])
	}
}

func ReduceBinary(f Binary) Reduce {
	return func(ns ...int) int {
		return f(ns[0], ns[1])
	}
}
