package lineshift

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func FileToByteslice(filename string) []byte {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	byteSlice := make([]byte, size_of_slice)

	carriageReturn := ([]byte("\x0D"))

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("% X %c", byteSlice, byteSlice)
	fmt.Println()
	fmt.Println("Det er", bytes.Count(byteSlice, carriageReturn), "Carriage returns.")
	return byteSlice
}
