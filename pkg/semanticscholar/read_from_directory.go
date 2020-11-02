package semanticscholar

import (
	"log"
	"os"
	"path/filepath"
)

// visit walks over files in a directory
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("visit", err)
		}
		// do not include directories
		if !info.IsDir() {
			// only files
			*files = append(*files, path)
		}
		return nil
	}
}

// ReadFromDirectory parses the directory of separated files provided by semantic scholar
func ReadFromDirectory(directoryPath string) (results []*Publication) {
	log.Println("Start restoring directory:", directoryPath)

	var filPaths []string // stores the file paths of all the files in the directory

	// walk over the files in the directory
	err := filepath.Walk(directoryPath, visit(&filPaths))
	if err != nil {
		panic(err)
	}

	// Read all files
	for _, file := range filPaths {
		docs, err := ParseFile(file)
		if err != nil {
			log.Fatal("read file", err)
		}
		// add the parsed documents to the results
		results = append(results, docs...)
	}
	return
}
