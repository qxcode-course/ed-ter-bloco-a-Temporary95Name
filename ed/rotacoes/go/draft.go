package main
import "fmt"

func rotateSliceByN(slice []int) []int {

    // Fazer uma slice com capacidade "len(slice)" para armazenar os elementos rotacionados "n" vezes.
    sliceInvertida := make([]int, 0, len(slice))

    for i := len(slice) - 1; i >= 0; i-- {
		sliceInvertida = append(sliceInvertida, slice[i])
	}

    
    return sliceInvertida
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

    resultado := rotateSliceByN(sliceUm)
    

    // fmt.Println(a, b)
    // fmt.Println(sliceUm)

    fmt.Println(resultado)
}
