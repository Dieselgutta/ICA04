package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanLines()
	scanRunes()
	scanWords()
}
func scanLines() {
	file, err := os.Open("text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("Lines: %d\n", count)
}

func scanRunes() {
	// Open file and create scanner on top of it
	file, err := os.Open("text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanRunes)

	// Scan for next token.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("Runes: %d\n", count)

}

func scanWords() {
	// An artificial input source.
	file, err := os.Open("text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("Words: %d\n", count)

}
