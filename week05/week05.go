package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "cs r?cks~"
	replaceWords := strings.NewReplacer("?", "o")
	fixedWords := replaceWords.Replace(broken)
	fmt.Println(fixedWords)
}
