package elasticsearch

import (
	"context"
	"io/ioutil"
)

func Search(query string) ([]byte, error) {

	res, _ := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithQuery(query),
		es.Search.WithPretty(),
		es.Search.WithSourceExcludes("_class"),
	)

	return ioutil.ReadAll(res.Body)

	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())
}
