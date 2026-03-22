/*
package main

import (

	"fmt"

)

	func main() {
		var totalAlbum, qtdPossui int

		// Lendo as configurações iniciais
		if _, err := fmt.Scan(&totalAlbum); err != nil {
			return
		}
		if _, err := fmt.Scan(&qtdPossui); err != nil {
			return
		}

		possui := make([]int, qtdPossui)
		contagem := make(map[int]int)
		var repetidas []int

		// Lendo as figurinhas e identificando as repetidas
		for i := 0; i < qtdPossui; i++ {
			fmt.Scan(&possui[i])
			num := possui[i]

			// Se já vimos esse número antes, ele é uma repetida
			if contagem[num] > 0 {
				repetidas = append(repetidas, num)
			}
			contagem[num]++
		}

		// 1. Saída das Repetidas
		if len(repetidas) > 0 {
			for i, v := range repetidas {
				fmt.Print(v)
				if i < len(repetidas)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		} else {
			fmt.Println("N")
		}

		// 2. Identificando as Faltantes
		var faltantes []int
		for i := 1; i <= totalAlbum; i++ {
			// Se o número não está no nosso mapa de contagem, está faltando
			if contagem[i] == 0 {
				faltantes = append(faltantes, i)
			}
		}

		// Saída das Faltantes
		if len(faltantes) > 0 {
			for i, v := range faltantes {
				fmt.Print(v)
				if i < len(faltantes)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		} else {
			fmt.Println("N")
		}
	}
*/
package main

import (
	"fmt"
	"strings"
)

// 1. Função para identificar figurinhas repetidas
func buscarRepetidas(possuidas []int) []int {
	contagem := make(map[int]int)
	var repetidas []int

	for _, num := range possuidas {
		if contagem[num] > 0 {
			repetidas = append(repetidas, num)
		}
		contagem[num]++
	}
	return repetidas
}

// 2. Função para identificar quais estão faltando
func buscarFaltantes(totalAlbum int, possuidas []int) []int {
	// Transformamos em map para busca O(1)
	mapPossui := make(map[int]bool)
	for _, num := range possuidas {
		mapPossui[num] = true
	}

	var faltantes []int
	for i := 1; i <= totalAlbum; i++ {
		if !mapPossui[i] {
			faltantes = append(faltantes, i)
		}
	}
	return faltantes
}

// 3. Função auxiliar para formatar a saída conforme o requisito (N se vazio)
func formatarSaida(lista []int) string {
	if len(lista) == 0 {
		return "N"
	}
	// Converte []int para string separada por espaços
	return strings.Trim(fmt.Sprint(lista), "[]")
}

func main() {
	var totalAlbum, qtdPossui int

	// Entrada
	if _, err := fmt.Scan(&totalAlbum, &qtdPossui); err != nil {
		return
	}

	possuidas := make([]int, qtdPossui)
	for i := 0; i < qtdPossui; i++ {
		fmt.Scan(&possuidas[i])
	}

	// Processamento (Chamada das funções)
	repetidas := buscarRepetidas(possuidas)
	faltantes := buscarFaltantes(totalAlbum, possuidas)

	// Saída
	fmt.Println(formatarSaida(repetidas))
	fmt.Println(formatarSaida(faltantes))
}