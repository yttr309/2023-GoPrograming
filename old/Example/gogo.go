package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.8
	b := int(a)
	fmt.Println(math.Ceil(a))
	fmt.Println(math.Floor(a))
	fmt.Println(b)
}
