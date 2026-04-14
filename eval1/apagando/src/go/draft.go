package main

import "fmt"

func removeCommonSet(vet1, vet2 []int) []int {

	aux := make([]int, 0, len(vet1))

	for i := 0; i < len(vet1); i++ {
		for j := 0; j < len(vet2); j++ {
			if vet1[i] != vet2[j] {
				vet1[i] = vet2[j]
			} else {
				continue
			}
		}
	}

	return aux
}

func main() {
	var num1 int
	var num2 int

	fmt.Scanln(&num1)
	vet1 := make([]int, 0, num1)
	for i := 0; i < num1; i++ {
		fmt.Scanln(vet1[i])
	}

	fmt.Scanln(&num2)
	vet2 := make([]int, 0, num2)
	for i := 0; i < num2; i++ {
		fmt.Scanln(vet2[i])
	}

	resultado := removeCommonSet(vet1, vet2)
	fmt.Println(resultado)
}
