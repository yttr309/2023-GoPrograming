package main

import (
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	println(year)
	month := now.Month()
	println(month)
	//println(now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}
