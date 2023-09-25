package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	a := 1.2
	var b int = int(math.Floor(a))

	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	/* var a int = 9
	var b float32 = 2.7
	var c, d string = "inha", "go..."
	var f string
	var g bool
	h := 'z'
	i := "hi"
	j := [3]int{1, 2, 3}

	fmt.Println(a, b, c, d, f, g, h, j)

	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j)) */
	/*fmt.Println(reflect.TypeOf(('B')))
	fmt.Println(reflect.TypeOf((100)))
	fmt.Println(reflect.TypeOf((1.2)))
	fmt.Println(reflect.TypeOf((false)))
	fmt.Println(reflect.TypeOf(("Go!")))

	fmt.Println('가')
	fmt.Println("A")
	fmt.Println('김', '\n')
	fmt.Println(math.Floor(2.15))
	fmt.Println(math.Ceil(2.15))
	fmt.Println(strings.Title("오픈소스프로그래밍\n\"Go\""))
	fmt.Println(strings.Title("open source programing go!"))*/
}
