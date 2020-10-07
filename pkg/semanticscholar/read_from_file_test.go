package semanticscholar

import (
	"log"
	"testing"
)

func TestParseFile(t *testing.T) {
	// res, err := ParseFile("./testdata/test.jsonl")
	res, err := ParseFile("/Users/sebastianerhardt/Downloads/s2-corpus-000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Amount of results", len(res))
	log.Println("Example: ", res[10].Title)
}
