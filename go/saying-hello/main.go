package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("What is your name? ")

	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	name := strings.TrimSpace(input)

	output := "Hello, " + name + ", nice to meet you!"
	fmt.Println(output)
}
