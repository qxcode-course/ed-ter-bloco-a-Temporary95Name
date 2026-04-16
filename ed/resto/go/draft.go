package main

import "fmt"

func dividirPorDoisRecursao(n int) {

	if n == 0 {
		return
	}
	
	dividirPorDoisRecursao(n / 2)

	fmt.Println(n/2, n%2)
}

func main() {
	var numero int
	fmt.Scanln(&numero)

	dividirPorDoisRecursao(numero)
}
