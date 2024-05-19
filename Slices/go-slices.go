package main

import "fmt"

func main() {
	// Slices are an abstraction above arrays in Go. We can change size of slice but not of array.
	var myArr [3]string  // array initialize
	var mySlice []string // slice initialize

	myArr[0] = "3"
	myArr[1] = "4"
	myArr[2] = "5"

	mySlice = append(mySlice, "1")
	mySlice = append(mySlice, []string{"2", "3"}...)
	fmt.Println(mySlice)
	newSlic := make([]string, 3)
	copy(newSlic, mySlice)

	myArray := [3]string{"First", "Second", "Third"}

	mySlice = myArray[:]
	mySlice2 := myArray[:]

	mySlice[0] = "test"

	fmt.Println(mySlice2[0]) //"test"

	nSlice1 := []string{"1", "2"}
	nSlice2 := nSlice1[:]

	nSlice1[0] = "0"
	fmt.Println(nSlice1, nSlice2)
}
