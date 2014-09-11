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

//FactorOut returns the second argument after factoring out the first argument as many times as possible
func FactorOut(n, m int) int {
	if n == 0 {
		return m
	}
	for m%n == 0 {
		m /= n
	}
	return m
}

//FactorOutX returns a FactorOut Unary bound to a factor
func FactorOutX(n int) Unary {
	return Partial2(FactorOut, n)
}
