package main

import (
	"fmt"
	"go_sample/internal/sample"
)

func main() {
	// Run the main function of the app
	sample.HelloWorld()
	a := sample.Add(1, 2)
	fmt.Printf("1 + 2 = %d\n", a)
}
