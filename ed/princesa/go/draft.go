package main

import "fmt"

func print_jog(jogadores []int, espada int) {
    fmt.Print("[ ")
    for i, elem := range jogadores {
        if elem == 0 {
            continue
        }
        fmt.Print(elem)
        if i == espada {
            fmt.Print(">")
        }
        fmt.Print(" ")
    }
    fmt.Print("]\n")
}

func procurar_vivo(jogadores []int, espada int) int {
    for {
        espada = (espada + 1) % len(jogadores)
        if jogadores[espada] != 0 {
            return espada
        }
    }
}

func main () {
    var qtd, espada int
    fmt.Scan(&qtd, &espada)
    jogadores := make([]int, 0, qtd)

    for i := 1; i <= qtd; i++ {
        jogadores = append(jogadores, i)
    }
    espada -= 1
    for range qtd - 1 {
        print_jog(jogadores, espada)
        vai_morrer := procurar_vivo(jogadores, espada)
        jogadores[vai_morrer] = 0
        espada = procurar_vivo(jogadores, espada)
    }
    print_jog(jogadores, espada)
}