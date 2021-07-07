package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func calculater(n1 string, n2 string, s string) int {
	var x [2]int
	x[0], _ = strconv.Atoi(n1)
	x[1], _ = strconv.Atoi(n2)

	var result int

	if s == "+" {
		result = x[0] + x[1]
	} else if s == "-" {
		result = x[0] - x[1]
	} else if s == "*" {
		result = x[0] * x[1]
	}
	return result
}
*/
func search(ls []string) int {
	length := len(ls)
	for i := 0; i < length; i++ {
		if ls[i] == "+" {
			return i
		} else if ls[i] == "-" {
			return i
		} else if ls[i] == "*" {
			return i
		} else {
			return 2
		}
	}
	return 0
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	l := strings.Split(input, " ")

	//fmt.Println(calculater(l[0], l[1], l[2]))
	fmt.Println(len(l))

}

// 1 2 + 3 +
// 1 2 3 + + => 1 5 + >= 6
//記号が出てくるまで走査する？で、直前の二つ前をcalc？
