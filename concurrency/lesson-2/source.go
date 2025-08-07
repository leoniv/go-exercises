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

	pipe := fmap(join(
		producer(ctx, 1),
		producer(ctx, 2),
		producer(ctx, 3),
	),
		func(i int) string {
			return fmt.Sprintf("Resive:", i)
		},
	)

	result := fold(pipe, "", func(a, b string) string {
		return a + ", " + b
	})

	fmt.Println(result)
}

func fold[A, B any](in <-chan A, zero B, f func(A, B) B) B {
	result := zero
	drain(in, func(a A) {
		result = f(a, result)
	})
	return result
}

func drain[A any](in <-chan A, f func(A)) {
	for a := range in {
		f(a)
	}
}

func fmap[A, B any](in <-chan A, f func(A) B) <-chan B {
	result := make(chan B)

	go func() {
		defer close(result)
		for a := range in {
			result <- f(a)
		}
	}()

	return result
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

func producer(ctx context.Context, id int) <-chan int {
	ch := make(chan int)

	go func() {
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
	}()

	return ch
}
