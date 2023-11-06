package main

import (
	"fmt"
	"log"
	"os"
)

func isP(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("%d는 소수가 아닙니다", n)
	}
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false, nil
		}
	}
	return true, nil
}

func prime(n int) {
	p, err := isP(n)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if p {
		fmt.Println(n, "은 소수임")
	} else {
		fmt.Println(n, "은 소수가 아님")
	}
}
func primeRange(a int, b int) {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isP(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	/*
		fmt.Print("정수 입력 : ")
		rd := bufio.NewReader(os.Stdin)
		in, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		in = strings.TrimSpace(in)
		number, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal(err)
		}
	*/
	var menu int
	for true {
		fmt.Print("MENU : 1) 소수판정 2) 구간 소수판정 3) 종료 : ")
		_, err := fmt.Scanln(&menu)
		if err != nil {
			log.Fatal(err)
		}
		switch menu {
		case 1:
			fmt.Println("정수 입력 : ")
			var in int
			_, err := fmt.Scanln(&in)
			if err != nil {
				log.Fatal(err)
			}
			prime(in)
		case 2:
			fmt.Println("2가지 정수 입력 : ")
			var n1, n2 int
			_, err := fmt.Scanln(&n1, &n2)
			if err != nil {
				log.Fatal(err)
			}
			primeRange(n1, n2)
		default:
			fmt.Print("프로그램을 종료합니다.")
			os.Exit(1)
		}
	}

}
