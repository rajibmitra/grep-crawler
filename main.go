package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func main() {
	// Ensure the user provides the pattern and at least one directory argument
	if len(os.Args) < 3 {
		fmt.Println("Usage: grep-tool <pattern> <directory1> [directory2...]")
		os.Exit(1)
	}

	// Extract the pattern and directory arguments
	pattern := os.Args[1]
	directories := os.Args[2:]

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Iterate over the provided directories and search concurrently
	for _, dir := range directories {
		wg.Add(1)
		go func(directory string) {
			defer wg.Done()

			fmt.Printf("Searching for '%s' in '%s' (excluding .git directory):\n", pattern, directory)

			// Run the grep command with the pattern, directory, and exclude .git directory
			cmd := exec.Command("grep", "-rn", "--exclude-dir=.git", pattern, directory)

			// Capture the output of grep
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Error running grep: %v\n", err)
				return
			}

			// Print the grep results
			fmt.Printf("%s\n", output)
		}(dir)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
