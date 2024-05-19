package main

import "fmt"

func main() {
	var age int
	var name string
	var done bool
	const fixx = 10
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(done)

	s := "hello, 世界"
	for i, r := range s {
		fmt.Printf("%d -> %q\n", i, r)
	}

}
