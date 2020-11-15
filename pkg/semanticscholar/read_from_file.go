package semanticscholar

import (
	"bufio"
	"compress/gzip"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"path"
)

// use faster parser
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ParseFile takes a file name read the data from within the file
// and returns an array of parse Publications.
// It also checks if the file is in a compressed format like .gz
func ParseFile(fileName string) (results []*Publication, err error) {
	// init the read
	var scanner *bufio.Scanner
	// check if rawContent is compressed
	fileExtension := path.Ext(fileName)
	if fileExtension == ".gz" {
		// if file has a .gz ending
		compressedFile, errOpen := os.Open(fileName)
		if errOpen != nil {
			log.Println(errOpen)
			return nil, errOpen
		}
		defer compressedFile.Close()
		// get the raw content of the file
		rawContent, errGzip := gzip.NewReader(compressedFile)
		if errGzip != nil {
			log.Println(errGzip)
			return nil, errGzip
		}
		scanner = bufio.NewScanner(rawContent)
	} else {
		// if file has no file ending that indicates compression
		fileContent, errOpen := os.Open(fileName)
		if errOpen != nil {
			log.Println(errOpen)
			return nil, errOpen
		}
		defer fileContent.Close()
		// init scanner
		scanner = bufio.NewScanner(fileContent)
	}

	// create line buffer
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 30*1024*1024) // 300mb

	// iterate over the lines
	for scanner.Scan() {
		res, errLine := ParseLine(scanner.Bytes())
		if errLine != nil {
			log.Println(errLine)
			err = errLine
			return
		}
		results = append(results, &res)
	}

	err = scanner.Err()
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// ParseLine takes a line in byte from
// and returns a parse publication
func ParseLine(line []byte) (data Publication, err error) {
	err = json.Unmarshal(line, &data)
	return
}
