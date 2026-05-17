// equacao para a resposta e: (n elevado a 2) + (n vezes 2)
// nao sei se o codigo pode ser aceito ou completo ou correto.
package main

import "fmt"

func recPadrao2(n, originalN int) {

	if n < (originalN*originalN)+originalN+originalN {

		recPadrao2(n+originalN, originalN)

	} else {

		fmt.Printf("%d\n", n)
		return

	}

}

func main() {

	var n int

	fmt.Scanln(&n)

	// resultado := (n*n)+(n*2)

	// fmt.Printf("%d\n", (n*n)+(n*2))
	// fmt.Printf("\n")

	recPadrao2(n, n)
}
