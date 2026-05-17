package main

import (
	"fmt"
)

// Função recursiva que calcula n^2 e imprime o passo a passo
func calcularQuadrado(n int) int {
	// Caso Base
	if n == 1 {
		fmt.Println("1^2 = 1")
		return 1
	}

	// 1. Fase de Descida: Imprime a expressão com '?'
	fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)

	// Chamada recursiva para obter o valor de (n-1)^2
	quadradoAnterior := calcularQuadrado(n - 1)

	// Cálculo do valor atual baseado na fórmula de Aragão
	resultadoAtual := quadradoAnterior + 2*(n-1) + 1

	// 2. Fase de Subida: Imprime a expressão resolvida com o resultado
	fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, resultadoAtual)

	return resultadoAtual
}

func main() {
	var n int
	// Lendo a entrada do usuário
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	// Garante que o programa funcione corretamente para n >= 1
	if n >= 1 {
		calcularQuadrado(n)
	}
}