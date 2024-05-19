package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	nP := Person{
		Name: "dd",
		Age:  1,
	}
	fmt.Printf("%+v", nP)

	jsonOb, err := json.MarshalIndent(nP, " ", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonOb))

}
