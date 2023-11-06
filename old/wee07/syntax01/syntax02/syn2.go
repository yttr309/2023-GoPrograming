package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

}
