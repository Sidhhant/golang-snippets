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
	nP := Person{Name: "ee", Age: 123}

	// 1
	jsonOb, err := json.MarshalIndent(nP, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonOb))

	// 2
	fmt.Printf("%+v\n", nP)
}
