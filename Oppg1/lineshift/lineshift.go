package lineshift

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func FileToByteslice(filename string) []byte {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	lineBreak := ([]byte("\x0A"))

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("% X %c", byteSlice, byteSlice)
	fmt.Println()
	fmt.Println("Det er", bytes.Count(byteSlice, lineBreak), "linjeskift.")
	return byteSlice
}
