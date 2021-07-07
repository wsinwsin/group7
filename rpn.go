package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	l := strings.Split(input, " ")
	fmt.Println(1)

}
