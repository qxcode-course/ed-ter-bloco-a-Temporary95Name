// equacao para a resposta e: (n elevado a 2) + (n vezes 2)
// Ainda nao usa recursao
package main

import "fmt"

func recPadrao (n int) {

	// fmt.Printf("%d", n)
	if n == (n*n) + (n + n) {
		// fmt.Printf("%d", n)
		// return


		recPadrao(n + n)
	} else {
		// recPadrao(n + n)
		fmt.Printf("%d", n)
		return
	}
	// fmt.Printf("%d", n)
	// fmt.Println(n)
}



func main() {

	var n int

	fmt.Scanln(&n)

	// resultado := (n*n)+(n*2)

	// fmt.Printf("%d\n", (n*n)+(n*2))
	// fmt.Printf("\n")

	recPadrao(n)
}
