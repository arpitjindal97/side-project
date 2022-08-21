package docs

// internal/utils/swagger.go is imported automatically which contains
// all the responses

// swagger:route GET /api/torrents/{id} Torrent GetTorrentById
// Get a particular torrent by passing infohash
// responses:
//	200: torrent
//	default: error

// swagger:route DELETE /api/torrents/{id} Torrent DeleteTorrentById
// Deletes a torrent from Database
// responses:
//	200: message
//	default: error

// swagger:route GET /api/files/{id} Torrent GetFilesByInfohash
// Get all the files present inside a torrent
// responses:
//	200: files
//	default: error

// swagger:parameters GetTorrentById DeleteTorrentById GetFilesByInfohash
type infohashParam struct {
	// Pass the infohash id
	// in: path
	ID string `json:"id"`
}

// swagger:route POST /api/torrents Torrent PostTorrentById
// Send a request to add torrent
// responses:
//	200: message
//	default: error

// swagger:parameters PostTorrentById
type infohashPost struct {
	// infohash of torrent, one at a time
	// in: body
	PostBody struct {
		Infohash string `json:"infohash"`
	}
}

// swagger:route GET /search Torrent SearchQuery
// This powerful API provides all kinds of search functionalities with Lucene Syntax
// responses:
//	200: SearchQuery

// swagger:parameters SearchQuery
type searchQuery struct {
	// query parameter (lucene syntax) because backend is ElasticSearch
	// Ref:
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html#query-string-syntax
	// https://www.lucenetutorial.com/lucene-query-syntax.html
	// https://lucene.apache.org/core/2_9_4/queryparsersyntax.html
	// ex: name:(Rise) AND seeders:(>1)
	// in: query
	Query string `json:"q"`

	// Sort the results by one or more fields (lucene syntax)
	// Sorting to perform. Can either be in the form of fieldName, or fieldName:asc/fieldName:desc.
	// The fieldName has to be an actual field within the document
	// A comma-separated list of <field>:<direction> pairs
	// Ref: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	// ex: peers:desc,date:asc
	// in: query
	Sort string `json:"sort"`

	// Number of items per page
	// in: query
	Size int `json:"size"`

	// Tells from which item index result should appear (used together with size)
	// Index is the position of an element in result set (starting from zero)
	// It provides paging capabilities
	// -->
	// Ex: size=2 & from=2
	// ==> gives 1st page with 2 elements
	// Ex: size=2 & from=2
	// ==> gives 2nd Page with 2 elements
	// Ex: size=2 & from=1
	// ==> gives a page which has two elements but 1st element was also present at the bottom of previous page
	// in: query
	From int `json:"from"`
}
