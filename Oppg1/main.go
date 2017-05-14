package main

import (
	"flag"
	"log"

	"./fileutils"
	"./lineshift"
)

func main() {

	fileName := flag.String("f", "", "filename")

	flag.Parse()
	if *fileName == "" {
		log.Fatal("No filename specified, use -f")
	}

	fileutils.FileToByteslice(*fileName)
	lineshift.FileToByteslice(*fileName)

}
