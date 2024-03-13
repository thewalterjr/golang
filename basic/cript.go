package main

import (
	"fmt"
	"strconv"
)

func criptografarTelefone(telefone string) string {
	var telefoneCriptografado string
	for _, digito := range telefone {
		digitoCriptografado := (int(digito-'0') + 2) % 10
		telefoneCriptografado += strconv.Itoa(digitoCriptografado)
	}
	return telefoneCriptografado
}

func descriptografarTelefone(telefoneCriptografado string, chave string) string {
	if chave != "suco de laranja!🍊" {
		return "Chave incorreta!"
	}

	var telefoneDescriptografado string
	for _, digito := range telefoneCriptografado {
		digitoOriginal := (int(digito-'0') - 2 + 10) % 10
		telefoneDescriptografado += strconv.Itoa(digitoOriginal)
	}
	return telefoneDescriptografado
}

func main() {
	telefoneCriptografado := criptografarTelefone("5579996928345")
	fmt.Println("Telefone Criptografado:", telefoneCriptografado)

	telefoneDescriptografado := descriptografarTelefone(telefoneCriptografado, "valtindaDois2")
	fmt.Println("Telefone Descriptografado:", telefoneDescriptografado)
}
