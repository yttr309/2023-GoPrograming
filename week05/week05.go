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

	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScore = strings.TrimSpace(inputScore)
	score, err := strconv.ParseInt(inputScore, 32)
	if score >= 60 {
		grade := "A grade"
	} else {
		grade := "under A grade"
	}
	fmt.Println(inputScore)
	fmt.Println("You will get ", grade)

	//broken := "cs r?cks~"
	//replaceWords := strings.NewReplacer("?", "o")
	//fixedWords := replaceWords.Replace(broken)
	//fmt.Println(fixedWords)
}
