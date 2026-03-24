package main

import(
	"fmt"
	"strings"
)

func cleanInput(text string) []string{
	// Convert to lowercase
    lowered := strings.ToLower(text)
    
    // Trim leading and trailing whitespace
    trimmed := strings.TrimSpace(lowered)
    
    // Split on whitespace
    words := strings.Fields(trimmed)
    
    return words
}

func main() {
	fmt.Println("Hello, World!")
}