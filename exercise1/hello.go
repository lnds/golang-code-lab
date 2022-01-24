package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("type your name: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello ", text)
}
