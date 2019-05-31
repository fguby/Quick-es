package elastic

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

//列出所有文档
func ListMapping(index string) error {
	mapping := elastic.NewGetMappingService(client)
	docs, err := mapping.Index(index).Do(context.Background())
	if err != nil {
		return err
	}
	for _, v := range docs[index].(map[string]interface{}) {
		for k, v1 := range v.(map[string]interface{}) {
			fmt.Printf("索引下的文档为%v\n", k)
			fmt.Printf("对应的值为%v\n", v1)
		}
	}
	return nil
}
