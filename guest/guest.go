package main

import "fmt"

func main() {
	fmt.Println("Hello, wasmtime!")
	fmt.Printf("2 + 3 = %d\n", Add(2, 3))
}

//go:wasmexport Add
func Add(a, b int32) int32 {
	fmt.Println("Hello from Add()...")
	result := a + b
	fmt.Printf("Add(%d, %d) result is %d\n", a, b, result)
	return result
}

//go:wasmexport CallAdd
func CallAdd() {
	Add(40, 2)
}
