package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("숫자 입력 : ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)
	dan, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 10; i++ {
		fmt.Println(dan, "*", i, "=", (int(dan) * i))
	}
}
