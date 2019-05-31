package main

import (
	"errors"
	"log"
	"os"
	"quick-es/Workflow"
	"quick-es/elastic"

	"github.com/urfave/cli"
)

var ErrPathLack = errors.New("请指定配置文件路径")

func main() {
	var path string
	var index string
	app := cli.NewApp()
	app.Name = "quick es"
	app.Usage = "a quickly way for use elasticsearch"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "indexes, i",
			Usage:       "List index's informations，query all indexes by '*'",
			Destination: &index,
		},
		cli.StringFlag{
			Name:        "server, s",
			Usage:       "ElasticSearch's server address",
			Destination: &path,
		},
		//创建索引
		cli.StringFlag{
			Name:  "create, ci",
			Usage: "Create index by `indexname`",
		},
		//删除索引
		cli.StringFlag{
			Name:  "delete, di",
			Usage: "Create index by `indexname`",
		},
		//列出文档
		cli.StringFlag{
			Name:  "mapping, m",
			Usage: "Create index by `indexname`",
		},
	}
	app.Action = func(c *cli.Context) error {
		err := elastic.InitClient(path)
		if err != nil {
			wf := Workflow.New()
			wf.Add(false, "错误", "ES连接失败", "img/错误.icns", "")
			wf.SendFeedback()
			os.Exit(1)
		}
		// if c.NArg() > 0 {
		// 	name = c.Args().Get(0)
		// }
		//列出索引信息
		if index != "" {
			elastic.CatIndex(index)
		}
		//创建一个索引
		if c.String("create") != "" {
			elastic.CreateIndex(c.String("create"))
		}
		//删除一个索引
		if c.String("delete") != "" {
			elastic.DeleteIndex(c.String("delete"))
		}
		//列出type
		if c.String("mapping") != "" {
			elastic.ListMapping(c.String("mapping"))
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
