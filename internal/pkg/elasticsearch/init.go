package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v7"
)

var index string

var es *elasticsearch.Client

func Init(address []string, username, password, ind string) {

	esConfig := elasticsearch.Config{
		Addresses: address,
		Username:  username,
		Password:  password,
	}

	es, _ = elasticsearch.NewClient(esConfig)
	index = ind

}
