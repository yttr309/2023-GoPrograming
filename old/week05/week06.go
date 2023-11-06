package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	anwser := rand.Intn(100) + 1

	fmt.Println("Guess number (1 ~ 100):")
	fmt.Println(anwser)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("You chance : ", 10-i)
		fmt.Print("Guess number : ")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)

		if err != nil {
			log.Fatal(err)
		}
		if inputNumber < anwser {
			fmt.Println("Ur guess number is lower than answer.")
		} else if inputNumber > anwser {
			fmt.Println("Ur guess number is higher than answer.")
		} else {
			fmt.Println("맞춤")
			break
		}
	}
}
