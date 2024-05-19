package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	nP := Person{
		Name: "ee",
		Age:  123,
	}

	jsonOb, err := json.MarshalIndent(nP, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(jsonOb))

	// Wrapping error to identify the cause easily
	err = fmt.Errorf("Err is err", nil)
	if err != nil {
		fmt.Printf("DoAnotherThing failed: %s", err)
	}
}
