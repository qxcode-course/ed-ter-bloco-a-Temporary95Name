// Código feito por máquina - Dia 22 de Março de 2026
package main

import "fmt"

// maisProximo vai de "f" até "pontoFinal" e conta os passos
func maisProximo(pontoComeco, pontoFinal, direcao int) int {
	passos := 0
	atual := pontoComeco

	// Continuar o loop até chegar ao "pontoFinal"
	for atual != pontoFinal {
		atual = atual + direcao
		passos++

		// Fazer com que o loop fique "circular"
		if atual > 15 {
			atual = 0
		} else if atual < 0 {
			atual = 15
		}
	}
	return passos
}

func main() {
	var h, p, f, d int

	// Ler h, p, f, d
	fmt.Scan(&h, &p, &f, &d)

	// Calcular as distancia
	distanciaDeA := maisProximo(f, h, d)
	distanciaDeB := maisProximo(f, p, d)

	// Comparar os resultados e imprimir a resposta
	if distanciaDeA <= distanciaDeB {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
