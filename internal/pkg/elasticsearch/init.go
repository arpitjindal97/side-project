package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var index string

var es *elasticsearch.Client

func Init(address []string, username, password, ind string) {

	log.Println("Initializing ElasticSearch")
	esConfig := elasticsearch.Config{
		Addresses: address,
		Username:  username,
		Password:  password,
	}

	es, _ = elasticsearch.NewClient(esConfig)
	index = ind
	log.Println("ElasticSearch Session is ready")
}
