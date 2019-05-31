package elastic

import (
	"github.com/olivere/elastic"
)

var (
	client *elastic.Client
)

func InitClient(path string) error {
	var err error
	if path == "" {
		path = "http://localhost:9200"
	}
	client, err = elastic.NewClient(elastic.SetURL(path),
		elastic.SetSniff(false))
	return err
}
