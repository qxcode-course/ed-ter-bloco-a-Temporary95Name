package main

import "fmt"

func recMDC(num1 int, num2 int) int {

	if num1 == 0 {
		return num2
	}
	
	if num2 == 0 {
		return num1
	}
	
	return recMDC(num2, num1 % num2)
}


func main() {
	var num1 int
	var num2 int
	fmt.Scan(&num1)
    fmt.Scan(&num2)
	
	resultado := recMDC(num1, num2)
	fmt.Println(resultado)
	
	
}