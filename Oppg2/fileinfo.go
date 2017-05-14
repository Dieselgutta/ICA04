package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	fileInfo := flag.String("f", "", "filename")

	flag.Parse()
	if *fileInfo == "" {
		log.Fatal("No filename specified, use -f")
	}

	path, err := filepath.Abs(*fileInfo)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Stat(*fileInfo)
	if err != nil {
		log.Fatal(err)
	}

	b := float64(file.Size())
	kib := b / 1024
	mib := kib / 1024
	gib := mib / 1024
	fmt.Printf("Information about '%s':\n", path)
	fmt.Printf("Size: %.0fB, %fKiB, %fMiB, %.9fGiB\n", b, kib, mib, gib)
	mode := file.Mode()

	fmt.Println("Name : ", file.Name())
	fmt.Println("Size : ", file.Size())
	fmt.Println("Mode/permission : ", file.Mode())

	if mode&os.ModeSymlink != os.ModeSymlink {
		fmt.Println("File is not a symbolic link")
	} else {
		fmt.Println("File is a symbolic link")
	}

	if mode&os.ModeAppend != os.ModeAppend {
		fmt.Println("File is not append only")
	} else {
		fmt.Println("File is append only")
	}

	if mode&os.ModeDevice != os.ModeDevice {
		fmt.Println("Is not a device file")
	} else {
		fmt.Println("Is a device file")
	}

	if mode&os.ModeCharDevice != os.ModeCharDevice {
		fmt.Println("File is a char device file")
		fmt.Println("File is not a block device file")
	} else {
		fmt.Println("File is not a char device file")
		fmt.Println("FIle is a block device file")
	}

	fmt.Println("Is a directory? : ", file.IsDir())
	fmt.Println("Is a regular file? : ", file.Mode().IsRegular())
	fmt.Println("Unix permission bits? : ", file.Mode().Perm())
	fmt.Println("Permission in string : ", file.Mode().String())
}
