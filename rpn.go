package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculater(l []string) int {
	var x [2]int
	x[0], _ = strconv.Atoi(l[0])
	x[1], _ = strconv.Atoi(l[1])

	var result int

	if l[2] == "+" {
		result = x[0] + x[1]
	} else if l[2] == "-" {
		result = x[0] - x[1]
	} else if l[2] == "*" {
		result = x[0] * x[1]
	}
	return result
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	l := strings.Split(input, " ")

	fmt.Println(calculater(l))

}
