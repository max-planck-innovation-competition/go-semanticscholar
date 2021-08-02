package main

import (
	"flag"
	"github.com/max-planck-innovation-competition/go-semanticscholar/pkg/semanticscholar"
	"log"
)

func main() {
	// flags
	exportGz := flag.Bool("export-gz", false, "do you want to export compressed files")
	combined := flag.Bool("combined", false, "do you want to combine all files by type")
	// include exclude
	includePublications := flag.Bool("publications", true, "do you want to include the publications")
	includeAuthors := flag.Bool("authors", true, "do you want to include the Authors")
	includeFieldOfStudies := flag.Bool("fieldOfStudies", true, "do you want to include the FieldOfStudies")
	includeAuthorPublicationEdges := flag.Bool("authorPublicationEdges", true, "do you want to include the AuthorPublicationEdges")
	includePublicationFieldOfStudyEdges := flag.Bool("publicationFieldOfStudyEdges", true, "do you want to include the PublicationFieldOfStudyEdges")
	includeInCitationEdges := flag.Bool("inCitationEdges", true, "do you want to include the InCitationEdges")
	includeOutCitationEdges := flag.Bool("outCitationEdges", true, "do you want to include the OutCitationEdges")
	// directory
	importDirectory := flag.String("import-directory", "./", "the directory with the files you want to transform")
	exportDirectory := flag.String("export-directory", "./", "the export directory for the transformed files")
	// transform data
	// transform data
	e := semanticscholar.ETL{
		ImportDirectory:                     *importDirectory,
		ExportDirectory:                     *exportDirectory,
		Compress:                            *exportGz,
		Combined:                            *combined,
		AddHeaders:                          true,
		IncludePublications:                 *includePublications,
		IncludeAuthors:                      *includeAuthors,
		IncludeFieldOfStudies:               *includeFieldOfStudies,
		IncludeAuthorPublicationEdges:       *includeAuthorPublicationEdges,
		IncludePublicationFieldOfStudyEdges: *includePublicationFieldOfStudyEdges,
		IncludeInCitationEdges:              *includeInCitationEdges,
		IncludeOutCitationEdges:             *includeOutCitationEdges,
	}
	err := e.TransformDirectory()
	if err != nil {
		log.Fatal(err)
	}
}
