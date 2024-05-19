package main

import "fmt"

func main() {
	// Slices in Go are strictly typed, meaning all elements in a slice must be of the same type. This is a contrast to Python lists, which can contain elements of different types.
	mixed := []interface{}{1, "hello", 3.14, true}
	for _, v := range mixed {
		fmt.Println(v) // The values will be printed out.
	}

	for _, v := range mixed {
		switch v := v.(type) {
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		case float64:
			fmt.Println("Float64:", v)
		case bool:
			fmt.Println("Boolean:", v)
		default:
			fmt.Println("Unknown type")
		}
	}
}
