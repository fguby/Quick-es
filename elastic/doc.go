package elastic

import (
	"context"
	"github.com/olivere/elastic"
	"quick-es/Workflow"
)

//列出所有文档
func ListMapping(index string) error {
	mapping := elastic.NewGetMappingService(client)
	docs, err := mapping.Index(index).Do(context.Background())
	if err != nil {
		return err
	}
	wf := Workflow.New()
	for _, v := range docs[index].(map[string]interface{}) {
		for k, _ := range v.(map[string]interface{}) {
			wf.Add(false, k, "......", "img/1.icns", "")
		}
	}
	wf.SendFeedback()
	return nil
}

//创建一个文档
func CreateType(index, doc string) error {
	mapping := elastic.NewPutMappingService(client)
	_, err := mapping.Index(index).Type(doc).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
