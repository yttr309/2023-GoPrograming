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
	score, err := strconv.ParseFloat(inputScore, 64)
	var grade string
	if score >= 90.0 {
		grade = "A grade"
	} else {
		grade = "under A grade"
	}
	fmt.Println(inputScore)
	fmt.Println("You will get ", grade)

	//broken := "cs r?cks~"
	//replaceWords := strings.NewReplacer("?", "o")
	//fixedWords := replaceWords.Replace(broken)
	//fmt.Println(fixedWords)
}
