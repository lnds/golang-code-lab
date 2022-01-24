package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// strings are composed of runes
	vowels := make(map[rune]int)


	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ToLower(text)

		// iterate over every rune of text
		for _, c := range text {
			switch c {
			case 'a', 'e', 'i', 'o', 'u':
				vowels[c]++
			}
		}
	}

	fmt.Println("statistics:")
	for k, v := range vowels {
		fmt.Printf("vowel %c appears %d times\n", k, v)
	}
	
}