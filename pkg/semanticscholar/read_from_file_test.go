package semanticscholar

import (
	"log"
	"testing"
)

// Test std file
func TestParseFile(t *testing.T) {
	results, err := ParseFile("../../test/s2-corpus-000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Amount of results:", len(results))
	log.Println("Example:", results[0].Title)
}

// Tests gz file
func TestParseFileGz(t *testing.T) {
	results, err := ParseFile("../../test/s2-corpus-000.gz")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Amount of results:", len(results))
	log.Println("Example:", results[10].Title)
}
