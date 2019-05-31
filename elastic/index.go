package elastic

import (
	"context"
	"quick-es/Workflow"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
)

type MyIndex struct {
	Health    string
	Status    string
	DocsCount int
	Pri       int
	Rep       int
}

//索引一个新文档
//也可单独创建索引
func CreateIndex(index string) error {
	_, err := client.CreateIndex(index).Do(context.Background())
	return err
}

//删除索引
func DeleteIndex(index string) error {
	_, err := client.DeleteIndex(index).Do(context.Background())
	return err
}

//列出所有索引
func CatIndex(index string) error {
	cat := elastic.NewCatIndicesService(client)
	if index != "*" {
		index += "*"
	}
	indices, err := cat.Index(index).Do(context.Background())
	if err != nil {
		return err
	}
	wf := Workflow.New()
	if len(indices) == 0 {
		wf.Add(false, "什么都没有", "换个索引名称试试吧", "img/sorry.png", "")
	}
	//循环展示
	myIndex := &MyIndex{}
	for i, value := range indices {
		myIndex.Health = value.Health
		myIndex.DocsCount = value.DocsCount
		myIndex.Status = value.Status
		myIndex.Pri = value.Pri
		myIndex.Rep = value.Rep
		wf.Add(true, value.Index, myIndex.formart(), "img/"+strconv.Itoa(i+1)+".icns", value.Index)
	}
	wf.SendFeedback()
	return nil
}

//格式化输出
func (m *MyIndex) formart() string {
	var build strings.Builder
	build.WriteString("Health: ")
	build.WriteString(m.Health)
	build.WriteString(" | Status: ")
	build.WriteString(m.Status)
	build.WriteString(" | Pri: ")
	build.WriteString(strconv.Itoa(m.Pri))
	build.WriteString(" | Rep: ")
	build.WriteString(strconv.Itoa(m.Rep))
	build.WriteString(" | ")
	build.WriteString("DocsCount: ")
	build.WriteString(strconv.Itoa(m.DocsCount))
	return build.String()
}

//
