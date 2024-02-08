package main

import "fmt"

func sequencia(binario string) int {
	aux := 0

	for i := 0; i < len(binario)-2; i++ {
		if binario[i:i+3] == "100" {
			aux++
		}
	}

	return aux
}

func main() {
	fmt.Println("Resp", sequencia("10110010111010010010000"))
}
