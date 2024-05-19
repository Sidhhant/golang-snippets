package main

import "fmt"

func main() {
	// Format Specifiers
	/*
		s, p, d, v
		%q - single quoted character
	*/

	str := "helo, into the wild"

	for i, c := range str {
		// fmt.Printf("%v, %v", i, c)
		fmt.Printf("%d, %q, %d, %#v, %d\n", i, c, c, c, c)
	}

}
