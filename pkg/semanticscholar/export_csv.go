package semanticscholar

import (
	"compress/gzip"
	"encoding/csv"
	"os"
	"strconv"
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
	"name",
}
var publicationNodesHeader = []string{
	"publicationId:ID(Publication-ID)",
	"title",
	"paperAbstract",
	"s2url",
	"sources",
	"pdfUrls",
	"year:int",
	"venue",
	"journalName",
	"journalVolume",
	"journalPages",
	"doi",
	"doiUrl",
	"pmId",
	"magId",
}
var fieldOfStudyNodesHeader = []string{
	"fieldOfStudyId:ID(Field-Of-Study-ID)",
}
var author2PublicationEdgesHeader = []string{":START_ID(Author-ID)", ":END_ID(Publication-ID)", ":TYPE"}
var publication2FieldsOfStudyEdgesHeader = []string{":START_ID(Publication-ID)", ":END_ID(Field-Of-Study-ID)", ":TYPE"}
var publication2publicationEdgesHeader = []string{":START_ID(Publication-ID)", ":END_ID(Publication-ID)", ":TYPE"}

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
	publication2publicationEdges [][]string, // publication -> publication
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
		publication2publicationEdges = append(publication2publicationEdges, publication2publicationEdgesHeader)
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
			CleanString(pub.Title),
			CleanString(pub.PaperAbstract),
			CleanString(pub.S2URL),
			CleanString(strings.Join(pub.Sources, " | ")),
			CleanString(strings.Join(pub.PdfUrls, " | ")),
			CleanString(strconv.Itoa(pub.Year)),
			CleanString(pub.Venue),
			CleanString(pub.JournalName),
			CleanString(pub.JournalVolume),
			CleanString(pub.JournalPages),
			CleanString(pub.Doi),
			CleanString(pub.DoiURL),
			CleanString(pub.PmID),
			CleanString(pub.MagID),
		})

		// iterate over authors
		for _, a := range pub.Authors {
			for _, id := range a.IDs {
				// add author
				authors[id] = CleanString(a.Name)
				// add edge author to publication
				author2PublicationEdges = append(author2PublicationEdges, []string{id, pub.ID, "AUTHOR_OF"})
			}
		}

		// iterate over fields of study
		for _, f := range pub.FieldsOfStudy {
			fieldsOfStudy[f] = true
			publication2FieldsOfStudyEdges = append(publication2FieldsOfStudyEdges, []string{pub.ID, CleanString(f), "FIELDS_OF_STUDY"})
		}

	}

	for _, pub := range pubs {
		// publication 2 publication
		for _, p := range pub.InCitations {
			// only add the publication if its in the index
			// if publications[p] && publications[pub.ID] {
			publication2publicationEdges = append(publication2publicationEdges, []string{p, pub.ID, "CITES"})
			//}
		}
		for _, p := range pub.OutCitations {
			// only add the publication if its in the index
			// if publications[p] && publications[pub.ID] {
			publication2publicationEdges = append(publication2publicationEdges, []string{pub.ID, p, "CITES"})
			//}
		}
	}

	for k, _ := range fieldsOfStudy {
		fieldsOfStudyNodes = append(fieldsOfStudyNodes, []string{k})
	}

	for k, v := range authors {
		authorNodes = append(authorNodes, []string{k, v})
	}

	return
}

// ExportCsv transforms the data and stores it in a (compressed) csv file
func ExportCsv(gzip, addHeaders bool, onlyHeaders bool, publications []*Publication, exportFolderPath, prefix, suffix string) (err error) {
	authorNodes,
		publicationNodes,
		fieldsOfStudyNodes,
		author2PublicationEdges,
		publication2FieldsOfStudyEdges,
		publication2publicationEdges := generateRecords(addHeaders, onlyHeaders, publications)
	// author nodes
	err = WriteFile(gzip, authorNodes, exportFolderPath+"/"+prefix+"author_nodes"+suffix)
	if err != nil {
		return
	}
	// publication nodes
	err = WriteFile(gzip, publicationNodes, exportFolderPath+"/"+prefix+"publication_nodes"+suffix)
	if err != nil {
		return
	}
	// fields of study
	err = WriteFile(gzip, fieldsOfStudyNodes, exportFolderPath+"/"+prefix+"fieldsOfStudy_nodes"+suffix)
	if err != nil {
		return
	}
	// author to publication edges
	err = WriteFile(gzip, author2PublicationEdges, exportFolderPath+"/"+prefix+"author2Publication_edges"+suffix)
	if err != nil {
		return
	}
	// publication to field of study edges
	err = WriteFile(gzip, publication2FieldsOfStudyEdges, exportFolderPath+"/"+prefix+"publication2FieldsOfStudy_edges"+suffix)
	if err != nil {
		return
	}
	// publication to publication edges
	err = WriteFile(gzip, publication2publicationEdges, exportFolderPath+"/"+prefix+"publication2publication_edges"+suffix)
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