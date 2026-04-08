/*
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func piramide(slice []int) {
	
	if len(slice) < 1 {
		return
	}

	if len(slice) == 1 {
		fmt.Println(slice[0])
		return
	}

	// Create the next level: sum n and n+1
	nextLevel := make([]int, len(slice)-1)
	for i := 0; i < len(slice)-1; i++ {
		nextLevel[i] = slice[i] + slice[i+1]
	}

	piramide(nextLevel)

	for i, val := range slice {
		fmt.Print(val)
		if i < len(slice) - 1 {
			fmt.Print(" ")
		}
	}
	fmt.Print()
}

// CollectIntegers reads from an io.Reader and returns a slice of integers.
// By using io.Reader instead of *os.File, this function can read from
// files, network connections, or even strings during testing.
func CollectIntegers(r io.Reader) ([]int, error) {
	// Initialize an empty slice. It starts with length 0.
	var result []int

	// Create a scanner to wrap our reader.
	// Scanner is efficient for reading stream data.
	scanner := bufio.NewScanner(r)

	// CRITICAL STEP: Tell the scanner to split by "words" (spaces/newlines).
	// Without this, it would try to read whole lines as single integers.
	scanner.Split(bufio.ScanWords)

	// scanner.Scan() returns true as long as there is another "word" to read.
	// It returns false when it hits EOF (Ctrl+D) or an error.
	for scanner.Scan() {
		// scanner.Text() gives us the string representation of the current token.
		word := scanner.Text()

		// Attempt to convert the string (e.g., "123") into a base-10 integer.
		// strconv.Atoi is shorthand for ParseInt(s, 10, 0).
		val, err := strconv.Atoi(word)
		if err != nil {
			// If the user types "hello" instead of "123", we return the error.
			// You could also 'continue' here if you want to ignore bad input.
			return nil, fmt.Errorf("invalid input '%s': %w", word, err)
		}

		// Add the valid integer to our slice. 
		// Go handles the memory reallocation automatically as the slice grows.
		result = append(result, val)
	}

	// After the loop, check if Scan stopped because of a real error (like a disk failure)
	// rather than just reaching the end of the input.
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	// fmt.Println("Enter numbers (Type a non-number or press Ctrl+D on Unix/macOS or Ctrl+Z on Windows to finish):")

	// Call our function, passing in os.Stdin.
	// Because os.Stdin implements io.Reader, it fits perfectly.
	nums, err := CollectIntegers(os.Stdin)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		// We can still print what we managed to collect before the error
	}

	// fmt.Printf("%v\n", nums)
	piramide(nums)
}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumPyramid(arr []int) {
	if len(arr) == 0 {
		return
	}

	// Base Case: Top of the pyramid
	if len(arr) == 1 {
		printFormatted(arr)
		return
	}

	// Calculate the next level
	nextLevel := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		nextLevel[i] = arr[i] + arr[i+1]
	}

	// Recursive call (Dive deep first)
	sumPyramid(nextLevel)

	// Print current level during the "unwind"
	printFormatted(arr)
}

// Helper function to print with specific "[ 1 2 3 ]" spacing
func printFormatted(arr []int) {
	fmt.Print("[ ")
	for i, val := range arr {
		fmt.Print(val)
		if i < len(arr) {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

func main() {
	// fmt.Println("Enter integers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var inputArr []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err == nil {
				inputArr = append(inputArr, num)
			}
		}

		if len(inputArr) > 0 {
			sumPyramid(inputArr)
		}
	}
}