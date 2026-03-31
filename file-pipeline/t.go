package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Open the file
	file, err := os.Open("large_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// 2. Ensure file is closed
	defer file.Close()

	// 3. Create a scanner to read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 4. Process each line here (scanner.Text() or scanner.Bytes())
		fmt.Println(scanner.Text())
	}

	// 5. Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
