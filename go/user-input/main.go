package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	return input
}

func main() {
	var input string

	fmt.Println("What is the game name?")
	input = readInput()
	fmt.Printf("You entered: %q\n", input)

	fmt.Println("What is the game's Age Rating?")
	input = readInput()
	fmt.Printf("You entered: %q\n", input)
}
