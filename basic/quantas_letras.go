package main

import (
	"fmt"
	"strings"
)

func counterLetter(word, letter string) int {
	word = strings.TrimSpace(word)
	count := 0
	for _, char := range word {
		if string(char) == letter {
			count++
		}
	}

	return count

}

func counterLetter2(word, letter string) int {
	return strings.Count(word, letter)
}

func main() {
	fmt.Println("res", counterLetter2("macaco", "a"))
}
