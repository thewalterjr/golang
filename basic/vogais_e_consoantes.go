package main

import (
	"fmt"
	"strings"
)

func separarVogaisEConsoantes(word string) (string, string) {
	vogais := ""
	consoantes := ""

	for _, w := range word {
		if strings.ContainsRune("aeiou", w) {
			vogais += string(w)
		} else {
			consoantes += string(w)
		}
	}

	return vogais, consoantes
}

func main() {
	vogais, consoantes := separarVogaisEConsoantes("banana")

	fmt.Printf("Vogais: %s\n", vogais)
	fmt.Printf("Consoantes: %s\n", consoantes)
}
