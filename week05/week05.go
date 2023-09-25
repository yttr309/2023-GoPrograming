package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n')
	fmt.Println(inputScore)

	//broken := "cs r?cks~"
	//replaceWords := strings.NewReplacer("?", "o")
	//fixedWords := replaceWords.Replace(broken)
	//fmt.Println(fixedWords)
}
