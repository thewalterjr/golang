package main

import (
	"fmt"
)

func contarDias(lista []int) int {
	aux := 0
	total := 0

	for i, v := range lista {
		dia := i + 1
		total += v
		if total >= 1000000 && aux == 0 {
			aux = dia
		}

	}

	return aux
}

func main() {
	l := []int{100, 99900, 400000, 500000, 600000} // 4
	l1 := []int{1000000}                           // 1
	l2 := []int{520511, 373475, 566377, 906090, 435915, 101456, 368890, 24832, 64713} // 3

	fmt.Println("Res:", contarDias(l))
	fmt.Println("Res:", contarDias(l1))
	fmt.Println("Res:", contarDias(l2))
}
