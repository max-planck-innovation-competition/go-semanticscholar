package semanticscholar

import (
	"compress/gzip"
	"encoding/csv"
	"os"
	"strings"
)

/*
https://neo4j.com/docs/operations-manual/current/tools/import/file-header-format/#import-tool-header-format-header-files
The header file of each data source specifies how the data fields should be interpreted.
You must use the same delimiter for the header file and for the data files.

The header contains information for each field, with the format <name>:<field_type>.
The <name> is used for properties and node IDs. In all other cases, the <name> part of the field is ignored.
*/

var authorNodesHeader = []string{
	"authorId:ID(Author-ID)",
	// "name",
}
var publicationNodesHeader = []string{
	"publicationId:ID(Publication-ID)",
	// "title",
	// "paperAbstract",
	// "s2url",
	// "sources",
	// "pdfUrls",
	// "year:int",
	// "venue",
	// "journalName",
	// "journalVolume",
	// "journalPages",
	// "doi",
	// "doiUrl",
	// "pmId",
	// "magId",
}
var fieldOfStudyNodesHeader = []string{
	"fieldOfStudyId:ID(Field-Of-Study-ID)",
}
var author2PublicationEdgesHeader = []string{":START_ID(Author-ID)", ":END_ID(Publication-ID)", ":TYPE"}
var publication2FieldsOfStudyEdgesHeader = []string{":START_ID(Publication-ID)", ":END_ID(Field-Of-Study-ID)", ":TYPE"}
var inCitationEdgesHeader = []string{":START_ID(Publication-ID)", ":END_ID(Publication-ID)", ":TYPE"}
var outCitationEdgesHeader = []string{":START_ID(Publication-ID)", ":END_ID(Publication-ID)", ":TYPE"}

// CleanString repairs artifacts that are in the dataset
// e.g. German umlauts
func CleanString(dirty string) string {
	dirty = strings.ReplaceAll(dirty, `\"u`, `ü`)
	dirty = strings.ReplaceAll(dirty, `\"o`, `ö`)
	dirty = strings.ReplaceAll(dirty, `\"a`, `ä`)
	dirty = strings.ReplaceAll(dirty, `\"`, `"`)
	dirty = strings.ReplaceAll(dirty, `\"`, `"`)
	dirty = strings.ReplaceAll(dirty, `\`, `/`)
	dirty = strings.ReplaceAll(dirty, `""`, `"`)
	return dirty
}

// generateRecords transforms the Publication objects into the csv format
func generateRecords(addHeaders bool, onlyHeaders bool, pubs []*Publication) (
	authorNodes [][]string,
	publicationNodes [][]string,
	fieldsOfStudyNodes [][]string,
	author2PublicationEdges [][]string,
	publication2FieldsOfStudyEdges [][]string,
	inCitationEdges [][]string, // publication -> publication
	outCitationEdges [][]string, // publication -> publication
) {

	authors := map[string]string{}     // creates a map of authors and ids
	publications := map[string]bool{}  // creates a map of all publication ids
	fieldsOfStudy := map[string]bool{} // creates a map of all fields of study

	if addHeaders {
		// add headers
		authorNodes = append(publicationNodes, authorNodesHeader)
		publicationNodes = append(publicationNodes, publicationNodesHeader)
		fieldsOfStudyNodes = append(fieldsOfStudyNodes, fieldOfStudyNodesHeader)
		author2PublicationEdges = append(author2PublicationEdges, author2PublicationEdgesHeader)
		publication2FieldsOfStudyEdges = append(publication2FieldsOfStudyEdges, publication2FieldsOfStudyEdgesHeader)
		inCitationEdges = append(inCitationEdges, inCitationEdgesHeader)
		outCitationEdges = append(outCitationEdges, outCitationEdgesHeader)
		// if you are interested in only the headers
		if onlyHeaders {
			return
		}
	}

	for _, pub := range pubs {

		publications[pub.ID] = true

		// add publication
		publicationNodes = append(publicationNodes, []string{
			pub.ID,
			//CleanString(pub.Title),
			//CleanString(pub.PaperAbstract),
			//CleanString(pub.S2URL),
			//CleanString(strings.Join(pub.Sources, " | ")),
			//CleanString(strings.Join(pub.PdfUrls, " | ")),
			//CleanString(strconv.Itoa(pub.Year)),
			//CleanString(pub.Venue),
			//CleanString(pub.JournalName),
			//CleanString(pub.JournalVolume),
			//CleanString(pub.JournalPages),
			//CleanString(pub.Doi),
			//CleanString(pub.DoiURL),
			//CleanString(pub.PmID),
			//CleanString(pub.MagID),
		})

		// iterate over authors
		for _, a := range pub.Authors {
			for _, id := range a.IDs {
				// add author
				// authors[id] = CleanString(a.Name)
				// add edge author to publication
				author2PublicationEdges = append(author2PublicationEdges, []string{id, pub.ID, "AUTHOR_OF"})
			}
		}

		// iterate over fields of study
		for _, f := range pub.FieldsOfStudy {
			fieldsOfStudy[f] = true
			publication2FieldsOfStudyEdges = append(publication2FieldsOfStudyEdges, []string{pub.ID, CleanString(f), "FIELDS_OF_STUDY"})
		}

		// publication 2 publication

		// in citations
		for _, p := range pub.InCitations {
			// List of paper IDs which cited this paper.
			inCitationEdges = append(inCitationEdges, []string{pub.ID, p, "CITED_BY"})
		}
		// out citations
		for _, p := range pub.OutCitations {
			// List of IDs which this paper cited.
			outCitationEdges = append(outCitationEdges, []string{pub.ID, p, "CITES"})
		}
	}

	for k, _ := range fieldsOfStudy {
		fieldsOfStudyNodes = append(fieldsOfStudyNodes, []string{k})
	}

	for k, _ := range authors {
		//authorNodes = append(authorNodes, []string{k, v})
		authorNodes = append(authorNodes, []string{k})
	}

	return
}

// ExportCsv transforms the data and stores it in a (compressed) csv file
func ExportCsv(i int, gzip, addHeaders bool, onlyHeaders bool, publications []*Publication, exportFolderPath, prefix, suffix string) (err error) {
	authorNodes,
		publicationNodes,
		fieldsOfStudyNodes,
		author2PublicationEdges,
		publication2FieldsOfStudyEdges,
		inCitationEdges,
		outCitationEdges := generateRecords(addHeaders, onlyHeaders, publications)
	// author nodes
	err = WriteFile(gzip, authorNodes, exportFolderPath+"/"+prefix+"author-nodes"+suffix)
	if err != nil {
		return
	}
	if i == 0 {
		// fields of study
		err = WriteFile(gzip, fieldsOfStudyNodes, exportFolderPath+"/"+prefix+"fields-of-study-nodes"+suffix)
		if err != nil {
			return
		}
	}
	// publication nodes
	err = WriteFile(gzip, publicationNodes, exportFolderPath+"/"+prefix+"publication-nodes"+suffix)
	if err != nil {
		return
	}
	// author to publication edges
	err = WriteFile(gzip, author2PublicationEdges, exportFolderPath+"/"+prefix+"author-2-publication-edges"+suffix)
	if err != nil {
		return
	}
	// publication to field of study edges
	err = WriteFile(gzip, publication2FieldsOfStudyEdges, exportFolderPath+"/"+prefix+"publication-2-fields-of-study-edges"+suffix)
	if err != nil {
		return
	}
	// in citations
	err = WriteFile(gzip, inCitationEdges, exportFolderPath+"/"+prefix+"in-citation-edges"+suffix)
	if err != nil {
		return
	}
	// out CitationEdges
	err = WriteFile(gzip, outCitationEdges, exportFolderPath+"/"+prefix+"out-citation-edges"+suffix)
	if err != nil {
		return
	}
	return
}

// ExportCsv transforms the data and stores it in a (compressed) csv file
func ExportAppendCsv(i int, publications []*Publication, exportFolderPath, prefix, suffix string) (err error) {
	authorNodes,
		publicationNodes,
		fieldsOfStudyNodes,
		author2PublicationEdges,
		publication2FieldsOfStudyEdges,
		inCitationEdges,
		outCitationEdges := generateRecords(false, false, publications)
	// author nodes
	err = AppendFile(authorNodes, exportFolderPath+"/"+prefix+"author-nodes"+suffix)
	if err != nil {
		return
	}
	if i == 0 {
		// fields of study
		err = AppendFile(fieldsOfStudyNodes, exportFolderPath+"/"+prefix+"fields-of-study-nodes"+suffix)
		if err != nil {
			return
		}
	}
	// publication nodes
	err = AppendFile(publicationNodes, exportFolderPath+"/"+prefix+"publication-nodes"+suffix)
	if err != nil {
		return
	}

	// author to publication edges
	err = AppendFile(author2PublicationEdges, exportFolderPath+"/"+prefix+"author-2-publication-edges"+suffix)
	if err != nil {
		return
	}
	// publication to field of study edges
	err = AppendFile(publication2FieldsOfStudyEdges, exportFolderPath+"/"+prefix+"publication-2-fields-of-study-edges"+suffix)
	if err != nil {
		return
	}
	// in citations
	err = AppendFile(inCitationEdges, exportFolderPath+"/"+prefix+"in-citation-edges"+suffix)
	if err != nil {
		return
	}
	// out CitationEdges
	err = AppendFile(outCitationEdges, exportFolderPath+"/"+prefix+"out-citation-edges"+suffix)
	if err != nil {
		return
	}
	return
}

func WriteFile(gzip bool, data [][]string, filePath string) (err error) {
	if gzip {
		return writeCSVGz(data, filePath)
	} else {
		return writeCSV(data, filePath)
	}
}

// writeCSVGz generates a compressed csv file
func writeCSVGz(data [][]string, filePath string) (err error) {
	file, err := os.Create(filePath + ".csv.gz")
	if err != nil {
		return
	}
	// init writers
	gzipWriter := gzip.NewWriter(file)
	csvWriter := csv.NewWriter(gzipWriter)
	// write the data
	err = csvWriter.WriteAll(data)
	if err != nil {
		return
	}
	csvWriter.Flush()
	err = gzipWriter.Flush()
	if err != nil {
		return
	}
	// close gzip writer
	err = gzipWriter.Close()
	if err != nil {
		return
	}
	// close file
	err = file.Close()
	if err != nil {
		return
	}
	return
}

// writeCSV generates a csv file
func writeCSV(data [][]string, filePath string) (err error) {
	// create file
	file, err := os.Create(filePath + ".csv")
	if err != nil {
		return
	}
	// create writer
	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(data)
	if err != nil {
		return
	}
	// close file
	err = file.Close()
	if err != nil {
		return
	}
	return
}

// AppendFile appends the content to all file
func AppendFile(data [][]string, filePath string) (err error) {
	file, err := os.OpenFile(filePath+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	// create writer
	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(data)
	if err != nil {
		return
	}
	// close file
	err = file.Close()
	if err != nil {
		return
	}
	return
}
