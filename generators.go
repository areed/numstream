package numstream

func ToSlice(c <-chan int) []int {
	var s []int
	for n := range c {
		s = append(s, n)
	}
	return s
}

func ToChannel(ns []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range ns {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func Generate(next Reduce, proceed Predicate, seed ...int) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for _, n = range seed {
			ch <- n
		}
		state := seed
		for proceed(n) {
			n = next(state...)
			state = append(state[1:], n)
			ch <- n
		}
		close(ch)
	}()
	return ch
}
