package main

import (
	"fmt"
)

// Formata a saída da lista de jogadores conforme o estado atual
func printJogadores(jogadores []int, indiceEspada int) {
	fmt.Print("[ ")
	for i, val := range jogadores {
		if val > 0 {
			if i == indiceEspada {
				fmt.Printf("%d> ", val)
			} else {
				fmt.Printf("%d ", val)
			}
		} else {
			if i == indiceEspada {
				fmt.Printf("<%d ", val)
			} else {
				fmt.Printf("%d ", val)
			}
		}
	}
	fmt.Println("]")
}

// Encontra o próximo jogador vivo em uma direção específica (1 para direita, -1 para esquerda)
func buscarProximoVivo(jogadores []int, atual int, direcao int) int {
	n := len(jogadores)
	pos := (atual + direcao + n) % n
	for {
		if jogadores[pos] != 0 {
			return pos
		}
		pos = (pos + direcao + n) % n
	}
}

func main() {
	var n, e, f int
	if _, err := fmt.Scan(&n, &e, &f); err != nil {
		return
	}

	// Inicializa os jogadores com sinais alternados
	jogadores := make([]int, n)
	for i := 0; i < n; i++ {
		// Se f=1:  1, -2,  3, -4...
		// Se f=-1: -1,  2, -3,  4...
		val := i + 1
		if i%2 != 0 {
			jogadores[i] = -val * f
		} else {
			jogadores[i] = val * f
		}
	}

	// O índice da espada é E-1 (ajuste para 0-based)
	posEspada := e - 1

	for len(jogadores) > 1 {
		printJogadores(jogadores, posEspada)

		// Determina direção da morte baseada no sinal de quem está com a espada
		direcaoMorte := 1
		if jogadores[posEspada] < 0 {
			direcaoMorte = -1
		}

		// Identifica a vítima e remove do slice
		indiceVitima := buscarProximoVivo(jogadores, posEspada, direcaoMorte)
		
		// Criar novo slice sem a vítima para manter a "fila circular" compacta
		// e facilitar a impressão correta dos testes
		jogadores = append(jogadores[:indiceVitima], jogadores[indiceVitima+1:]...)

		// Ajusta a posição da espada após a remoção
		// Se a vítima estava antes da espada, o índice da espada diminui
		if indiceVitima < posEspada {
			posEspada--
		}

		// A espada passa para o próximo sobrevivente na mesma direção
		posEspada = buscarProximoVivo(jogadores, posEspada, direcaoMorte)
	}

	// Imprime o vencedor final
	printJogadores(jogadores, 0)
}