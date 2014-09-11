package numstream

//Sum is an n-ary addition operation.
func Sum(ns ...int) int {
	l := len(ns)
	if l == 1 {
		return ns[0]
	}
	return ns[0] + Sum(ns[1:]...)
}

//Add is a binary addition operation.
func Add(m, n int) int {
	return m + n
}

//IncrementBy returns a Unary addition function.
func IncrementByX(n int) Unary {
	return Partial2(Add, n)
}

//Increment is a Reduce incrementer for use in generators
var Increment = ReduceUnary(IncrementByX(1))
