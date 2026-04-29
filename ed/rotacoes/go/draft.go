// A resposta tem que ter um espaco antes do primeiro elemento e um espaco depois to ultimo elemento.
// Exemplo: [*1*2*3*4*5*]

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func revertSlice(slice []int, b int) []int {

	// Fazer uma slice com capacidade "len(slice)" para armazenar os elementos rotacionados "n" vezes.
	sliceInvertida := make([]int, 0, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		sliceInvertida = append(sliceInvertida, slice[i])
	}

	return sliceInvertida
}

func invertFirstK(slice []int, b int) []int {

	b = b % len(slice)

	storeInversion := make([]int, 0, len(slice))

	for i := 0; i < b; i++ {
		storeInversion = append(storeInversion, slice[i])
	}

	return storeInversion
}

func appendAfterK(slice []int, b int) []int {

	b = b % len(slice)

	storeAfterK := make([]int, 0, len(slice))

	for i := 0; i < len(slice); i++ {

		if i >= b {
			storeAfterK = append(storeAfterK, slice[i])
		}

	}

	return storeAfterK
}

func main() {

	var a int
	var b int

	// Ler os dois inteiros.
	fmt.Scanln(&a, &b)

	// Fazer uma slice com capacidade "a", e depois ler os elementos e coloca-los em "sliceUm".
	sliceUm := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&sliceUm[i])
	}

	primeiroPasso := revertSlice(sliceUm, b)

	segundoPasso := invertFirstK(primeiroPasso, b)

	primeiraMetadeInversa := revertSlice(segundoPasso, b)

	fazerSegundaMetade := appendAfterK(primeiroPasso, b)

	inverterSegundaMetade := revertSlice(fazerSegundaMetade, b)

	// fmt.Println(primeiroPasso)

	// fmt.Print(primeiraMetadeInversa)
	// fmt.Println(fazerSegundaMetade)
	// fmt.Print(inverterSegundaMetade)

	sliceFinal := make([]int, 0, a)
	sliceFinal = append(primeiraMetadeInversa, inverterSegundaMetade...)

	// fmt.Print(resultadoFinal)


	// Imprimir com espaco antes do primeiro elemento e depois depois do ultimo elemento. // IA/LLM foi usado.

	strSliceFinal := make([]string, len(sliceFinal))
	for i, v := range sliceFinal {
		strSliceFinal[i] = strconv.Itoa(v)
	}

	resultado :=  "[ " + strings.Join(strSliceFinal, " ") + " ]"
	fmt.Println(resultado)
}