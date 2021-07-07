package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func calculater(n1 string, n2 string, s string) string {
	var x [2]int
	x[0], _ = strconv.Atoi(n1)
	x[1], _ = strconv.Atoi(n2)

	var result int
	var returnresult string

	if s == "+" {
		result = x[0] + x[1]
	} else if s == "-" {
		result = x[0] - x[1]
	} else if s == "*" {
		result = x[0] * x[1]
	}

	returnresult = strconv.Itoa(result)
	return returnresult
}

func search(ls []string) int {
	length := len(ls)
	for i := 0; i < length; i++ {
		if ls[i] == "+" {
			return i
		} else if ls[i] == "-" {
			return i
		} else if ls[i] == "*" {
			return i
		}
	}
	return 0
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	l := strings.Split(input, " ")

	for {
		i := search(l)
		if i == 0 {
			break
		}
		l[i-2] = calculater(l[i-2], l[i-1], l[i])
		l = remove(l, i)
		l = remove(l, i-1)
	}

	answer, _ := strconv.Atoi(l[0])
	fmt.Println(answer)
}

// 1 2 + 3 +
// 1 2 3 + + => 1 5 + >= 6
//記号が出てくるまで走査する？で、直前の二つ前をcalc？
