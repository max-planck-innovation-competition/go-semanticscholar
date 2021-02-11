package main

import (
	"flag"
	"github.com/max-planck-innovation-competition/go-semanticscholar/pkg/semanticscholar"
	"log"
)

func main() {
	// flags
	exportGz := flag.Bool("export-gz", false, "do you want to export compressed files")
	combined := flag.Bool("combined", false, "do you want to combine all files by type")
	importDirectory := flag.String("import-directory", "./", "the directory with the files you want to transform")
	exportDirectory := flag.String("export-directory", "./", "the export directory for the transformed files")
	// transform data
	err := semanticscholar.TransformDirectory(*importDirectory, *exportDirectory, *exportGz, *combined)
	if err != nil {
		log.Fatal(err)
	}
}
