package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inputScore)

	//broken := "cs r?cks~"
	//replaceWords := strings.NewReplacer("?", "o")
	//fixedWords := replaceWords.Replace(broken)
	//fmt.Println(fixedWords)
}
