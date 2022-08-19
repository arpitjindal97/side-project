package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"example.com/m/internal/pkg/cassandra"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
)

func Search(query, sort string, size, from int) ([]byte, error) {

	res, _ := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithQuery(query),
		es.Search.WithSort(sort),
		es.Search.WithFrom(from),
		es.Search.WithSize(size),
		es.Search.WithPretty(),
		es.Search.WithSourceExcludes("_class"),
	)

	return ioutil.ReadAll(res.Body)

	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())
}

func Update(torrent cassandra.Torrent) {
	body, _ := json.Marshal(torrent)
	request := esapi.UpdateRequest{
		Index:      index,
		DocumentID: torrent.InfoHash,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, body))),
	}
	res, err := request.Do(context.Background(), es)
	if err != nil {
		_ = fmt.Errorf("update: request: %w", err)
	}
	defer res.Body.Close()
}

func Delete(id string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
	}
	_, err := req.Do(context.Background(), es)
	return err
}
