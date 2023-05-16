package main

import "fmt"

func calcu(lista map[string][]float64) map[string]float64 {
	media := make(map[string]float64)
	for aluno, valor := range lista {
		var soma float64
		for i := 0; i < len(valor); i++ {
			soma += valor[i]
		}
		resul := soma / float64(len(valor))
		media[aluno] = resul
	}
	return media
}

func main() {
	lista := calcu(map[string][]float64{
		"Kaynan":    []float64{7.8, 8.9, 9.9, 10.0, 7.5},
		"Sanderson": []float64{5.5, 6.6, 7.7, 10, 7.7},
		"Luisa":     []float64{8.7, 9.8, 9.9, 0.1, 5.7},
		"Pedro":     []float64{10, 9.9, 9.9, 10.0, 7.5},
		"Matheus":   []float64{7.8, 7.8, 3.1, 2.1, 1.2},
	})
	for aluno, media := range lista {
		fmt.Printf("%s: %.2f\n", aluno, media)
	}

}
