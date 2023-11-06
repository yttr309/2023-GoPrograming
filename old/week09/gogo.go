package main

import "fmt"

func main() {
	result(4.2, 3.0)
	result(5.2, 2.3)
}

func result(a float64, b float64) {
	fmt.Printf("%.2f ", (a*b)/10.0)
}
