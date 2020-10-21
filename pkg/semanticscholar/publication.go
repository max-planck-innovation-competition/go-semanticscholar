package semanticscholar

type Publication struct {
	ID            string `json:"id"`            // S2 generated research paper ID
	Title         string `json:"title"`         // Research paper title
	PaperAbstract string `json:"paperAbstract"` // Extracted abstract of the paper
	// Entities      []string `json:"entities"` // Extracted entities (deprecated on 2019-09-17)
	FieldsOfStudy []string `json:"fieldsOfStudy"` // Zero or more fields of study this paper addresses
	S2URL         string   `json:"s2Url"`         // URL to S2 research paper details page
	PdfUrls       []string `json:"pdfUrls"`       // URLs related to this PDF scraped from the web
	Authors       []struct {
		Name string   `json:"name"` // Name of the author
		IDs  []string `json:"ids"`  // S2ID of the author
	} `json:"authors"` // List of authors with an S2 generated author ID and name
	InCitations   []string `json:"inCitations"`   // List of S2 paper IDs which cited this paper
	OutCitations  []string `json:"outCitations"`  // List of S2 paper IDs which this paper cited
	Year          int      `json:"year"`          // Year this paper was published as integer
	Venue         string   `json:"venue"`         // Extracted publication venue for this paper
	JournalName   string   `json:"journalName"`   // Name of the journal that published this paper
	JournalVolume string   `json:"journalVolume"` // The volume of the journal where this paper was published
	JournalPages  string   `json:"journalPages"`  // The pages of the journal where this paper was published
	Sources       []string `json:"sources"`       // Identifies papers sourced from DBLP or Medline
	Doi           string   `json:"doi"`           // Digital Object Identifier registered at doi.org
	DoiURL        string   `json:"doiUrl"`        // DOI link for registered objects
	PmID          string   `json:"pmid"`          // Unique identifier used by PubMed
	MagID         string   `json:"magId"`         // Unique identifier used by Microsoft Academic Graph
}
