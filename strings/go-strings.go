package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "testName"
	fmt.Println(name[4:6])
	fmt.Println(len(name))

	// copy string
	var nameCopy = name[:]
	fmt.Println(nameCopy)
	// Print address
	fmt.Println(&name)
	fmt.Println(&nameCopy)

	/*
		Strings are immutable that is if you reassign a varaible that was used to assign some other variable it won't change the value of second variable.
	*/

	var a = "x"
	var b = a
	a = "y"
	fmt.Println(a, b)

	// Functionalities
	if strings.HasPrefix("test", "te") {
		fmt.Println("te is prefix of test")
	}

	var str = "hi can you be my valentine. I got a car parked outside the building. Let's go."
	var splitStr []string
	splitStr = append(splitStr, strings.Split(str, " ")...)
	fmt.Println(splitStr[4:6])

	s1 := "123"
	s2 := s1[:]
	r := []rune(s1)
	r[0] = 'a'
	modifiedS1 := string(r)
	fmt.Println(s1, s2, modifiedS1)
}
