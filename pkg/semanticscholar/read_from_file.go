package semanticscholar

import (
	"bufio"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
)

// use faster parser
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ParseFile(fileName string) (results []*Publication, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	// init scanner with buffer size of 1MB
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
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

func ParseLine(line []byte) (data Publication, err error) {
	err = json.Unmarshal(line, &data)
	return
}
