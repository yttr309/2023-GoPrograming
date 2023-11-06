package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Println(number, "는 소수다")
	} else {
		fmt.Println(number, "는 소수가 아니다")
	}
}
