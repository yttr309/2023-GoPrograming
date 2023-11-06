package main

import "fmt"

func main() {
	var t int
	var a int
	t, a = process(11, 22, 44)
	name := "홍길동"
	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", name, t, a)
}

func sayHello() {
	fmt.Println("Hello")
}

func process(kor int, eng int, math int) (int, int) {
	totals := kor + eng + math
	avg := totals / 3

	return totals, avg
}
