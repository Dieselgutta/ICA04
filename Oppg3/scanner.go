package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	fileInfo := os.Args[1]

	scanLines(fileInfo, "\n")
	scanRunes(fileInfo, "\n")
	scanWords(fileInfo, "\n")
	temp := check(fileInfo)
	freqRunes(temp)

}
func scanLines(fileInfo string, search string) int {
	file, err := os.Open(fileInfo)
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
	return count
}

func scanRunes(fileInfo string, search string) int {
	// Open file and create scanner on top of it
	file, err := os.Open(fileInfo)
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
	return count

}

func scanWords(fileInfo string, search string) int {
	// An artificial input source.
	file, err := os.Open(fileInfo)
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
	return count

}

func checkRune(filename string, search string) int {
	//Åpner filen
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 512*1024)
	//Teller linjeskift
	counter := 0
	//Søker på linjeskift
	searchFor := []byte(search)

	c, _ := file.Read(buffer)
	counter += bytes.Count(buffer[:c], searchFor)
	return counter
}

func check(filename string) map[int]string {
	m := make(map[int]string)

	// Runer for Ascii code valgt:
	for i := 0x20; i <= 0x7F; i++ {
		count := checkRune(filename, string(i))
		rune := string(i)
		m[count] = rune
	}
	fmt.Println()
	return m
}

func freqRunes(m map[int]string) {
	//preallocate the map size
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Print 5 største:
	fmt.Println("5 mest brukte runer:")
	value1 := keys[len(keys)-1]
	value2 := keys[len(keys)-2]
	value3 := keys[len(keys)-3]
	value4 := keys[len(keys)-4]
	value5 := keys[len(keys)-5]
	fmt.Println("Antall:", value1, "Rune:", m[value1])
	fmt.Println("Antall:", value2, "Rune:", m[value2])
	fmt.Println("Antall:", value3, "Rune:", m[value3])
	fmt.Println("Antall:", value4, "Rune:", m[value4])
	fmt.Println("Antall:", value5, "Rune:", m[value5])
}
