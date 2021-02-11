package semanticscholar

var AuthorNodesHeader = []string{
	"authorId",
	"name",
}

var PublicationNodesHeader = []string{
	"publicationId",
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

var FieldOfStudyNodesHeader = []string{
	"fieldOfStudyId",
}

var Author2PublicationEdgesHeader = []string{
	"authorId",
	"publicationId",
	"type",
}

var Publication2FieldsOfStudyEdgesHeader = []string{
	"publicationId",
	"fieldOfStudyId",
	"type",
}

var InCitationEdgesHeader = []string{
	"publicationIdStart",
	"publicationIdEnd",
	"type",
}

var OutCitationEdgesHeader = []string{
	"publicationIdStart",
	"publicationIdEnd",
	"type",
}
