package numstream

import (
	"sync"
)

func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				merged <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}
