package main

import (
	"fmt"
	"mymath64"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath64.Sqrt64(2))

	slice := []int{1, 2, 3, 4}
	fmt.Println(slice[0])

	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	fmt.Println(numbers)
}
