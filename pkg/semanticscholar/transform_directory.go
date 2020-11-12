package semanticscholar

import (
	"log"
	"path/filepath"
	"strconv"
)

func TransformDirectory(importDirectory, exportDirectory string, compress bool) (err error) {
	log.Println("Start transforming directory:", importDirectory)
	var filePaths []string // stores the file paths of all the files in the directory
	// walk over the files in the directory
	err = filepath.Walk(importDirectory, visit(&filePaths))
	if err != nil {
		return err
	}
	// Read all files
	for i, file := range filePaths {
		log.Println("Process:", i, "/", len(filePaths))
		publications, errFile := ParseFile(file)
		if errFile != nil {
			log.Println("error while reading file: ", file, " : ", errFile)
			return errFile
		}
		// create header files
		// only do that once
		if i == 0 {
			errHeader := ExportCsv(compress, true, true, publications, exportDirectory, "", "-headers")
			if errHeader != nil {
				log.Println("error while exporting header files ", errHeader)
				return errHeader
			}
		}
		// create suffix with index
		suffix := "-data-" + strconv.Itoa(i)
		// export files
		errExport := ExportCsv(compress, false, false, publications, exportDirectory, "", suffix)
		if errExport != nil {
			log.Println("error while exporting files: ", errExport)
			return errExport
		}
	}
	log.Println("Done transforming directory:", importDirectory)
	log.Println("Exported to:", exportDirectory)
	return
}