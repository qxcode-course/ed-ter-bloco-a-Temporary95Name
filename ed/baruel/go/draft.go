// Resposta feita por máquina - 21 de Março de 2026
package main

import (
	"fmt"     // Package for formatted I/O (Input/Output)
	"strconv" // Package for conversions to/from string
	"strings" // Package for string manipulation like Join
)

// FindDuplicates extracts repeated values from the collection
func FindDuplicates(stickers []int) []int {
	counts := make(map[int]int) // Map to track how many times each ID appears
	var duplicates []int        // Slice to store the results

	for _, s := range stickers { // Loop through every sticker owned
		if counts[s] > 0 { // If ID was already seen at least once...
			duplicates = append(duplicates, s) // ...add it to the duplicates list
		}
		counts[s]++ // Increment the count for this sticker ID
	}
	return duplicates // Return the list of extra stickers
}

// FindMissing identifies gaps in the collection based on album size
func FindMissing(total int, owned []int) []int {
	ownedMap := make(map[int]bool) // Map for O(1) time complexity lookups
	for _, s := range owned {      // Iterate through the owned stickers
		ownedMap[s] = true // Mark this sticker ID as "possessed"
	}

	var missing []int                // Slice to store IDs not found
	for i := 1; i <= total; i++ {    // Iterate from 1 up to the album total
		if !ownedMap[i] { // If the ID is NOT in our "possessed" map...
			missing = append(missing, i) // ...add it to the missing list
		}
	}
	return missing // Return all IDs needed to finish the album
}

// FormatOutput converts a slice to a string or returns "N"
func FormatOutput(list []int) string {
	if len(list) == 0 { // Check if the slice is empty
		return "N" // Return "N" as per the problem requirement
	}

	strValues := make([]string, len(list)) // Create a string slice of the same size
	for i, v := range list {               // Loop through the integer slice
		strValues[i] = strconv.Itoa(v) // Convert each integer to a string
	}
	return strings.Join(strValues, " ") // Merge strings with a space between them
}

func main() {
	var totalAlbum, countOwned int // Declare variables for the initial metadata

	// Read the total album capacity and the number of stickers Baruel has
	if _, err := fmt.Scan(&totalAlbum, &countOwned); err != nil {
		return // Exit if there is an error reading the input
	}

	owned := make([]int, countOwned) // Initialize a slice for the owned stickers
	for i := 0; i < countOwned; i++ { // Loop for the amount of stickers owned
		fmt.Scan(&owned[i]) // Read each sticker ID into the slice
	}

	duplicates := FindDuplicates(owned)      // Call function to find repeats
	missing := FindMissing(totalAlbum, owned) // Call function to find gaps

	fmt.Println(FormatOutput(duplicates)) // Print the formatted duplicates line
	fmt.Println(FormatOutput(missing))    // Print the formatted missing line
}