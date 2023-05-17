package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Print("What is the input string? ")

	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	str := strings.TrimSuffix(input, "\n")
	numChars := strconv.Itoa(utf8.RuneCountInString(str))

	//fmt.Println("\"" + str + "\"")

	output := "\"" + str + "\" has " + numChars + " characters."
	fmt.Println(output)
}
