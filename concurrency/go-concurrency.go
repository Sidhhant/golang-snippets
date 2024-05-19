package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("stating here in go rountine")
		m <- 2
		fmt.Println("after sending 2 in m")
		m <- 3
		fmt.Println("after sending 3 in m")
		m <- 4
		fmt.Println("after sending 4 in m")
		fmt.Println("unlocked")
		fmt.Println("ending go rountine")
		wg.Done()
	}()
	fmt.Println("in main")
	fmt.Println(<-m)
	fmt.Println("in main after printing 2")
	fmt.Println(<-m)
	fmt.Println("in main after printing 3")
	fmt.Println(<-m)
	fmt.Println("in main after printing 4")
	wg.Wait()
}
