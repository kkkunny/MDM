package util

import "sync"

func MixRChannel[T any](ch ...<-chan T) <-chan T {
	out := make(chan T)

	var wg sync.WaitGroup
	for _, c := range ch {
		wg.Go(func() {
			for v := range c {
				out <- v
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
