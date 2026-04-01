package main

import (
	"fmt"
	"strings"
)

func countOccurencesInSlice (mainArray []string, queryArray []string) []int {
	// Mapa para fazer a contagem
	contagem := make(map[string]int)
	
	// Slice para armazenar os resultados
	resultado := make([]int, len(queryArray))
	
	
	for _, str := range mainArray {
		contagem[str]++
	}
	
	
	for i, consulta := range queryArray {
		resultado[i] = contagem[consulta]
	}
	
	
	return resultado
}


func main() {
	
    var mainSize int
    fmt.Scan(&mainSize) // Reads the first line integer

    mainArray := make([]string, mainSize)
    for i := 0; i < mainSize; i++ {
        fmt.Scan(&mainArray[i]) // Reads each string separated by space/newline
    }


    var querySize int
    fmt.Scan(&querySize) // Reads the third line integer

    queryArray := make([]string, querySize)
    for i := 0; i < querySize; i++ {
        fmt.Scan(&queryArray[i]) // Reads each query string
    }


    resultadoA1 := countOccurencesInSlice(mainArray, queryArray)
	
	result1 := fmt.Sprint(resultadoA1)
    result1 = strings.Trim(result1, "[]")
	
	fmt.Println(result1)
    /*
    sliceA1 := []string{"aba", "baba", "aba", "xzxb"}
	sliceA2 := []string{"aba", "xzxb", "ab"}
	
	sliceB1 := []string{"c", "c", "mpikv"}
	sliceB2 := []string{"f", "c", "o", "uplbd", "o", "zl", "xoi", "mpikv"}
	
	sliceC1 := []string{"qu", "qu", "pg", "qu", "qu", "cjyrx", "y"}
	sliceC2 := []string{"o", "qu", "sxsse", "b", "pd", "yuv"}
	
	resultadoA1 := countOccurencesInSlice(sliceA1, sliceA2)
	resultadoA2 := countOccurencesInSlice(sliceB1, sliceB2)
	resultadoA3 := countOccurencesInSlice(sliceC1, sliceC2)
	
	
	result1 := fmt.Sprint(resultadoA1)
    result1 = strings.Trim(result1, "[]")
    
    result2 := fmt.Sprint(resultadoA2)
    result2 = strings.Trim(result2, "[]")
    
    result3 := fmt.Sprint(resultadoA3)
    result3 = strings.Trim(result3, "[]")
	
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
    */
}
