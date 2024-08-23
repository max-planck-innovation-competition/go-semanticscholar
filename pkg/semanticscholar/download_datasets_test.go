package semanticscholar

import (
	"log"
	"testing"
)

// Test std file
func TestDownload(t *testing.T) {

	var ApiKey = "SetYour-SemanticScholar-APIKey"

	err := Download("../sample", ApiKey)
	if err != nil {
		log.Fatal(err)
	}

}
