package lesson1

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	dict := make(map[int]int)
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	for i := range 1000 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Millisecond)
			mu.RLock()
			defer mu.RUnlock()
			if v, ok := dict[i]; ok {
				fmt.Println("k:", i, "v:", v)
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 1000 {
			mu.Lock()
			dict[i] = i
			mu.Unlock()
		}
	}()

	wg.Wait()
}
