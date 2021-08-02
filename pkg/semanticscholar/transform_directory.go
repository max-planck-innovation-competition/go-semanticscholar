package semanticscholar

import (
	"log"
	"path/filepath"
	"strconv"
)

func (e *ETL) TransformDirectory() (err error) {
	log.Println("Start transforming directory:", e.ImportDirectory)
	var filePaths []string // stores the file paths of all the files in the directory
	// walk over the files in the directory
	err = filepath.Walk(e.ImportDirectory, visit(&filePaths))
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
			errHeader := e.ExportCsv(i, e.Compress, true, true, publications, "", "-headers")
			if errHeader != nil {
				log.Println("error while exporting header files ", errHeader)
				return errHeader
			}
		}
		// if the mode is Combined
		// the output is one file containing all of the
		if e.Combined {
			suffix := "-data-all"
			errExport := e.ExportAppendCsv(i, publications, "", suffix)
			if errExport != nil {
				log.Println("error while exporting files: ", errExport)
				return errExport
			}
		} else {
			// the output are multiple files
			suffix := "-data-" + strconv.Itoa(i)
			// export files
			errExport := e.ExportCsv(i, e.Compress, false, false, publications, "", suffix)
			if errExport != nil {
				log.Println("error while exporting files: ", errExport)
				return errExport
			}
		}
	}
	log.Println("Done transforming directory:", e.ImportDirectory)
	log.Println("Exported to:", e.ExportDirectory)
	return
}
