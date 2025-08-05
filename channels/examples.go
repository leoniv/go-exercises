package channels

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	input := make(chan int, 10)
	output := make(chan string, 10)
	var writerDone sync.WaitGroup
	readerDome := make(chan struct{})

	for id := range 3 {
		writerDone.Add(1)
		go func() {
			defer fmt.Println("Writer done:", id)
			defer writerDone.Done()
			writer(id, input, output)
		}()
	}

	go func() {
		defer fmt.Println("Reader done:")
		defer func() {
			readerDome <- struct{}{}
		}()

		reader(output)
	}()

	for msg := range 101 {
		input <- msg
	}
	close(input)
	writerDone.Wait()

	close(output)
	<-readerDome
}

func reader(output chan string) {
	for msg := range output {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func writer(id int, input chan int, output chan string) {
	for msg := range input {
		output <- fmt.Sprint("Worker: ", id, " message: ", msg)
		time.Sleep(10 * time.Millisecond)
	}
}
