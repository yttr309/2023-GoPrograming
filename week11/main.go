package main

import (
	"fmt"
)

func main() {

	primes := [3]int{2, 5, 1}

	primes[2] = 6

	test := [3]bool{true, false, true}

	fmt.Printf("%#v\n", test)

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}
}
