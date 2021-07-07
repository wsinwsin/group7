package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(l []string) int {
	var x [2]int
	x[0], _ = strconv.Atoi(l[0])
	x[1], _ = strconv.Atoi(l[1])

	result := x[0] + x[1]
	return result
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	l := strings.Split(input, " ")

	fmt.Println(add(l))

}
