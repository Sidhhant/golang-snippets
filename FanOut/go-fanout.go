package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Producer: ", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Processed: ", num, id)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(ch, &wg)

	// Fan-out
	for i := 0; i < 2; i++ {
		go consumer(ch, &wg, i)
	}

	wg.Wait()
}
