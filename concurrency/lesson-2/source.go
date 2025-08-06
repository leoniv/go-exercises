package lesson2

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run() {
	ctx1, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	ctx, _ := context.WithTimeout(ctx1, 3*time.Second)
	onExit := func() { fmt.Println("Program exit by: ", ctx.Err()) }
	defer onExit()

	chan1 := make(chan int, 1000)
	chan2 := make(chan int, 1000)
	go producer(ctx, chan1, 1)
	go producer(ctx, chan2, 2)

	out := join(chan1, chan2)

	for msg := range out {
		fmt.Println("Resive: ", msg)
	}
}

func join[T any](chans ...<-chan T) <-chan T {
	wg := sync.WaitGroup{}
	out := make(chan T)
	wg.Add(len(chans))

	for _, in := range chans {
		go func(in <-chan T) {
			defer wg.Done()

			for msg := range in {
				out <- msg
			}
		}(in)
	}

	go func() {
		defer fmt.Println("Exit join")
		defer close(out)
		wg.Wait()
	}()

	return out
}

func producer(ctx context.Context, ch chan<- int, id int) {
	defer fmt.Println("Producer interrupted")
	defer close(ch)

	ticker := time.NewTicker(time.Duration(id) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ch <- id
		}
	}
}
