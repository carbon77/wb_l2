package orchan

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	dones := make([]chan interface{}, len(channels))
	for i := range dones {
		dones[i] = make(chan interface{})
	}

	for i, channel := range channels {
		go func() {
			for {
				select {
				case <-channel:
					for _, done := range dones {
						close(done)
					}
					close(result)
					return
				case <-dones[i]:
					return
				}
			}
		}()
	}

	return result
}

func RunOrChannel() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v\n", time.Since(start))

	<-time.After(30 * time.Second)
}
