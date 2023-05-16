package main

import "fmt"

//Escreva uma função que receba um slice de inteiros e retorne um novo slice contendo
//apenas os valores únicos, ou seja, sem duplicatas. Utilize um mapa para auxiliar na
//remoção das duplicatas.

func atuslice(s []int) []int {
	var ns = []int{}
	mapa := make(map[int]int)
	for _, valor := range s {
		mapa[valor]++
	}
	for chaves := range mapa {
		ns = append(ns, chaves)
	}
	return ns
}

func main() {
	s := atuslice([]int{2, 2, 4, 8, 9, 9})
	fmt.Println(s)
}
