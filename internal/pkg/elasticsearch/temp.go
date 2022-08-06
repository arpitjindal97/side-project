package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func Search(query string) {
	esConfig := elasticsearch.Config{
		Addresses: []string{"http://vergon-elasticsearch:9200"},
		Username:  "elastic",
		Password:  "password",
	}

	es, _ := elasticsearch.NewClient(esConfig)
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
