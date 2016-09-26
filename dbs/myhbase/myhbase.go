package myhbase

import (
	"strings"
	"github.com/hunterhug/go-hbase"
	"errors"
)

//hbase配置结构体
/*
                "master": "192.168.11.73:60000",
                "zkport": "2181",
                "zkquorum": "192.168.11.73"
*/
type HbaseConfig struct {
	Master      string
	Zkport      string
	Zkquorum    string
}

type Hbase struct {
	Config HbaseConfig
	Client *hbase.Client
}

func New(config HbaseConfig) Hbase {
	return Hbase{Config:config}
}

//太坑了，坑啊，表名前缀不能太相似！！比如aaaaaaaaaaaaaaaaa,aaaaaaaaaaaaaaaaaab这样可能会把数据发送到另外的表
func (db *Hbase)Open() {

	config := db.Config
	zkhosts := strings.Split(config.Zkquorum, ",")
	for i, _ := range zkhosts {
		zkhosts[i] = zkhosts[i] + ":" + config.Zkport
	}
	zkroot := "/hbase"
	client := hbase.NewClient(zkhosts, zkroot)
	//client.SetLogLevel("DEBUG")
	db.Client = client
}

//获取结果
func (db *Hbase)GetResult(table string, rowkey string) (result *hbase.ResultRow, err error) {
	client := db.Client
	get := hbase.CreateNewGet([]byte(rowkey))
	result, err = client.Get(table, get)
	if (rowkey != result.Row.String()) {
		err = errors.New("没有rowkey")
	}
	return
}
