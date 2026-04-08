package main

import "fmt"

func eh_primo(numero, divisor int) bool {
	
	if numero <= 1 {
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

func main() {
	var num int
	
	fmt.Scanln(&num)
	
	var result bool = eh_primo(num, 2)
	
	fmt.Println(result)
}
