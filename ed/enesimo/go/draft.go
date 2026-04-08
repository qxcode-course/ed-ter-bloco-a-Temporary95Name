package main

import "fmt"

func eh_primo(numero, divisor int) bool {
	
	if numero <= 1 {
		fmt.Println("O numero e igual a 1.")
		return false
	}
	
	if divisor == numero {
		return true
	}
	
	if numero % divisor == 0 {
		return false
	}
	
	return eh_primo(numero, divisor + 1)
}

func enesimoPrimo(num1 int, num2 int, contagem int) int {
	
	if eh_primo(num2, 2) {
		contagem++
	}
	
	if contagem == num1 {
		return num2
	}
	
	return enesimoPrimo(num1, num2 + 1, contagem)
}

func main() {
	var num int
	
	fmt.Scanln(&num)
	
	result := enesimoPrimo(num, 2, 0)
	
	fmt.Println(result)
}
