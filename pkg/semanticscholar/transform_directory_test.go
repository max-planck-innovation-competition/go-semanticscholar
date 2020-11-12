package semanticscholar

import (
	"log"
	"testing"
)

// Tests the complete directory
func TestTransformDirectory(t *testing.T) {
	err := TransformDirectory("/media/seb/SCRAPER/semanticscholar", "/media/seb/SCRAPER/neo4j/", true)
	if err != nil {
		log.Fatal(err)
	}
}
