package semanticscholar

import (
	"log"
	"testing"
)

// Tests the complete directory
func TestTransformDirectory(t *testing.T) {
	e := ETL{
		ImportDirectory:                     "/media/seb/SCRAPER/semanticscholar",
		ExportDirectory:                     "/media/seb/SCRAPER/neo4j/import",
		Compress:                            true,
		Combined:                            true,
		AddHeaders:                          true,
		IncludePublications:                 true,
		IncludeAuthors:                      true,
		IncludeFieldOfStudies:               true,
		IncludeAuthorPublicationEdges:       true,
		IncludePublicationFieldOfStudyEdges: true,
		IncludeInCitationEdges:              true,
		IncludeOutCitationEdges:             true,
	}
	err := e.TransformDirectory()
	if err != nil {
		log.Fatal(err)
	}
}
