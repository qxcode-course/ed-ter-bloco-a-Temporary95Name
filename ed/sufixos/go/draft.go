package main
import "fmt"


func printRecursaoPiramide(texto string) {

    if len(texto) == 0 {
        return
    } 

    // Chamada recursiva
    printRecursaoPiramide(texto[1:len(texto)])
    
    // Armazenar o resultado na stack, para imprimir os resultados em ordem invertida
    fmt.Println(texto)

}


func main() {
    
    var texto string

    fmt.Scanln(&texto)

    printRecursaoPiramide(texto)
    
}
