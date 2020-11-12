package main

import (
	"flag"
	"github.com/Max-Planck-Innovation-Competition/go-semanticscholar/pkg/semanticscholar"
	"log"
)

func main() {
	// flags
	exportGz := flag.Bool("export-gz", false, "do you want to export compressed files")
	importDirectory := flag.String("import-directory", "./", "the directory with the files you want to transform")
	exportDirectory := flag.String("export-directory", "./", "the export directory for the transformed files")
	// transform data
	err := semanticscholar.TransformDirectory(*importDirectory, *exportDirectory, *exportGz)
	if err != nil {
		log.Fatal(err)
	}
}
