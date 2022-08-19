package utils

import (
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/pkg/elasticsearch"
)

// This file holds swagger models which are common to all microservices

// JSON Body of a Torrent
// swagger:response torrent
type torrentResponseWrapper struct {
	// in:body
	Torrent cassandra.Torrent
}

// JSON representation of all files present in a give torrent
// swagger:response files
type filesResponse struct {
	// in:body
	Files cassandra.Files
}

// Generic response containing a success message
// swagger:response message
type messageResponse struct {
	// in:body
	Message Message
}

// Generic response containing an error
// swagger:response error
type errorResponse struct {
	// in:body
	Error Error
}

// ElasticSearch Result JSON
// swagger:response SearchQuery
type elasticSearchResponse struct {
	// in:body
	SearchResponse elasticsearch.SearchResponse
}
