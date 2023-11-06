package main

import "fmt"

func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}
