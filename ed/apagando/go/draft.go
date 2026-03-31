package main

import (
	"fmt"
	"strings"
)

func removeFromSet(mainSet []int, toBeRemoved []int) []int {
	// Make map
	mapaRemoveElements := make(map[int]bool)
	for _, val := range toBeRemoved {
		mapaRemoveElements[val] = true
	}
	
	
	// Initialize slice and allocate capacity equal to mainSet slice
	resultado := make([]int, 0, len(mainSet))
	
	
	// Loop to check the boolean map
	for _, val := range mainSet {
		// If not in the map, the value becomes false
		if !mapaRemoveElements[val] {
			resultado = append(resultado, val)
		}
	}
	
	return resultado
}

func main() {
	// Objective: Make two slices with the size n1 and n2, and fill them with user defined integers.
	
	var n1 int // First slice size
	fmt.Scan(&n1)
	
	firstSet := make([]int, n1) // Create slice and store numbers in it
	for i := 0; i < n1; i++ {
		fmt.Scan(&firstSet[i]) 
	}
	
	
	var n2 int // Second slice size
	fmt.Scan(&n2)
	
	secondSet := make([]int, n2) // Create slice and store numbers in it
	for i := 0; i < n2; i++ {
		fmt.Scan(&secondSet[i]) 
	}
	
	
	intSlice := removeFromSet(firstSet, secondSet)
	
	// Convert slice to string
	tempString := fmt.Sprint(intSlice)
	// Remove "[" and "]" from the string 
	resultado := strings.Trim(tempString, "[]")
	
	// fmt.Println(firstSet)
	// fmt.Println(secondSet)
	fmt.Printf(resultado) 
	fmt.Print(" ")
	fmt.Println()
}