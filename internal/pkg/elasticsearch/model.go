package elasticsearch

import "example.com/m/internal/pkg/cassandra"

// https://github.com/elastic/go-elasticsearch/blob/main/_examples/encoding/model/response.go

type SearchResponse struct {
	Took int64 `json:"took"`
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*SearchHit `json:"hits"`
	} `json:"hits"`
}

type SearchHit struct {
	Score   float64 `json:"_score"`
	Index   string  `json:"_index"`
	Type    string  `json:"_type"`
	Version int64   `json:"_version,omitempty"`

	Source cassandra.Torrent `json:"_source"`
}
