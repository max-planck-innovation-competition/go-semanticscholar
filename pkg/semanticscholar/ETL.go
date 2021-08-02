package semanticscholar

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
}
