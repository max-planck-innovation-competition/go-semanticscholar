package semanticscholar

import (
	"strconv"
	"strings"
)

type ETL struct {
	ImportDirectory                     string
	ExportDirectory                     string
	Compress                            bool
	Combined                            bool
	AddHeaders                          bool
	IncludePublications                 bool
	IncludeAuthors                      bool
	IncludeFieldOfStudies               bool
	IncludeAuthorPublicationEdges       bool
	IncludePublicationFieldOfStudyEdges bool
	IncludeInCitationEdges              bool
	IncludeOutCitationEdges             bool
	PublicationFieldHandler             func(pub *Publication) []string
	PublicationHeaderHandler            func() []string
}

// CheckDefaultHandlers checks if there are handlers for the publications in place
// otherwise, use the default handlers
func (e *ETL) CheckDefaultHandlers() {
	if e.PublicationFieldHandler == nil {
		e.PublicationFieldHandler = DefaultPublicationFields
	}
	if e.PublicationHeaderHandler == nil {
		e.PublicationHeaderHandler = DefaultPublicationHeaders
	}
}

func (e *ETL) AddPublicationFieldHandler(fn func(pub *Publication) []string) *ETL {
	e.PublicationFieldHandler = fn
	return e
}

func (e *ETL) AddPublicationHeaderHandler(fn func() []string) *ETL {
	e.PublicationHeaderHandler = fn
	return e
}

func DefaultPublicationFields(pub *Publication) (result []string) {
	result = []string{
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
	}
	return
}

func DefaultPublicationHeaders() (result []string) {
	return PublicationNodesHeader
}
