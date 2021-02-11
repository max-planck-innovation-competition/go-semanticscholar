package semanticscholar

var Neo4jAuthorNodesHeader = []string{
	"authorId:ID(Author-ID)",
	"name",
}

var Neo4jPublicationNodesHeader = []string{
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

var Neo4jFieldOfStudyNodesHeader = []string{
	"fieldOfStudyId:ID(Field-Of-Study-ID)",
}

var Neo4jAuthor2PublicationEdgesHeader = []string{
	":START_ID(Author-ID)",
	":END_ID(Publication-ID)",
	":TYPE",
}

var Neo4jPublication2FieldsOfStudyEdgesHeader = []string{
	":START_ID(Publication-ID)",
	":END_ID(Field-Of-Study-ID)",
	":TYPE",
}
var Neo4jInCitationEdgesHeader = []string{
	":START_ID(Publication-ID)",
	":END_ID(Publication-ID)",
	":TYPE",
}

var Neo4jOutCitationEdgesHeader = []string{
	":START_ID(Publication-ID)",
	":END_ID(Publication-ID)",
	":TYPE",
}
