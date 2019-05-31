package test

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

const (
	//Version is the current version of Elastic
	//目前es版本
	Version = "6.2.17"

	//DefaultURL is the default endpoint of Elasticsearch on the local machine
	//Is is used e.g. when initializing a new client without a specific URL
	//本地es的路径
	//初始化一个没有特定URL的client时，用此值替代
	DefaultURL = "http://127.0.0.1:9200"

	//默认嗅探其余集群节点用的协议
	DefaultScheme = "http"

	//默认开启集群健康检查
	DefaultHealthcheckEnabled = true

	//默认健康检查的时长
	DefaultHealthcheckTimeoutStartup = 5 * time.Second

	//
	DefaultHealthcheckTimeout = 1 * time.Second
)

var (
	//没有可用的elastic节点
	ErrNoClient = errors.New("no Elasticsearch node available")

	//多次重试无法连接
	ErrRetry = errors.New("cannot connect after several retries")

	//请求超时将引发ErrTimeout错误
	ErrTimeout = errors.New("timeout")

	//
	//noRetries = NewStopRetrier()
)

//
type ClientOptionFunc func(*Client) error

type Client struct {
	c *http.Client

	connsMu sync.RWMutex
	//conns []*conn
}
