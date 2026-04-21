package main
import "fmt"

func rotateSliceByN(slice []int, n int) []int {

    // Fazer uma slice com capacidade "len(slice)" para armazenar os elementos rotacionados "n" vezes.
    rotacionados := make([]int, 0, len(slice))

    for 
    
    return rotacionados
}
func main() {

    var a int
    var b int

    // Ler os dois inteiros.
    fmt.Scanln(&a, &b)

    // Fazer uma slice com capacidade "a", e depois ler os elementos e coloca-los em "sliceUm". 
    sliceUm := make([]int, 0, a)
    for i := 0; i < a; i++ {
        fmt.Scanln(&sliceUm[i])  
    }

    resultado := rotateSliceByN(sliceUm, b)
    

    // fmt.Println(a, b)
    // fmt.Println(sliceUm)

    fmt.Println(resultado)
}
