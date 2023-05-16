package main

import (
	"fmt"
	"strings"
)

func lercaracter(s string) map[string]int {
	carac := strings.Split(s, "")
	ocorrencias := make(map[string]int)
	for _, cara := range carac {
		ocorrencias[cara]++
	}
	return ocorrencias
}

func main() {
	s := lercaracter("Yago vai tirar so ss")
	fmt.Println("A ocorrencia de cada caracter é: ", s)
}
