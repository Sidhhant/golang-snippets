package main

import "fmt"

func main() {
	// Println returns bytes written.
	bytes, err := fmt.Println("hello world")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bytes)
}
