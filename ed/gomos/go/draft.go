// Código feito por máquina - Dia 22 de Março de 2026
package main

import (
	"fmt"
)

// Ponto representa uma coordenada no plano 2D.
// No contexto deste problema:
// - O eixo x aumenta para a direita.
// - O eixo y aumenta para baixo (padrão de sistemas de coordenadas de tela).
type Ponto struct {
	X, Y int
}

func main() {
	var q int
	var direcao string

	// Lendo a configuração inicial:
	// Q = Quantidade de gomos
	// D = Direção (U, D, L, R)
	if _, err := fmt.Scan(&q, &direcao); err != nil {
		return
	}

	// Criamos um slice para armazenar as posições atuais de todos os gomos.
	// O gomo no índice 0 é sempre a cabeça da cobra.
	gomosAntigos := make([]Ponto, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&gomosAntigos[i].X, &gomosAntigos[i].Y)
	}

	// 1. CÁLCULO DA NOVA POSIÇÃO DA CABEÇA
	// A cabeça se move de forma independente baseada na direção fornecida.
	novaCabeca := gomosAntigos[0]
	switch direcao {
	case "U": // Up: Decrementa Y (sobe no plano)
		novaCabeca.Y--
	case "D": // Down: Incrementa Y (desce no plano)
		novaCabeca.Y++
	case "L": // Left: Decrementa X (vai para a esquerda)
		novaCabeca.X--
	case "R": // Right: Incrementa X (vai para a direita)
		novaCabeca.X++
	}

	// 2. SIMULAÇÃO DO MOVIMENTO DE RASTRO (CORPO)
	// A lógica fundamental:
	// - O gomo i na nova posição é igual ao gomo i-1 na posição anterior.
	// - Isso cria o efeito visual de que cada parte da cobra "segue" a anterior.
	novasPosicoes := make([]Ponto, q)
	novasPosicoes[0] = novaCabeca // A nova cabeça ocupa a posição calculada acima.

	for i := 1; i < q; i++ {
		// O gomo atual assume a posição que o gomo à sua frente ocupava antes do passo.
		novasPosicoes[i] = gomosAntigos[i-1]
	}

	// 3. SAÍDA DOS DADOS
	// Imprime cada coordenada (x y) em uma nova linha conforme solicitado.
	for _, p := range novasPosicoes {
		fmt.Printf("%d %d\n", p.X, p.Y)
	}
}