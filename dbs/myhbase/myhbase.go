package myhbase

import (
	"strings"
	"github.com/hunterhug/go-hbase"
	"clicki_web/conf"
	"errors"
)

//太坑了，坑啊，表名前缀不能太相似！！比如aaaaaaaaaaaaaaaaa,aaaaaaaaaaaaaaaaaab这样可能会把数据发送到另外的表
func GetInfoDb() *hbase.Client {

	config := conf.GetCenter().Resources.Hbase["info"]
	zkhosts := strings.Split(config.Zkquorum, ",")
	for i, _ := range zkhosts {
		zkhosts[i] = zkhosts[i] + ":" + config.Zkport
	}
	zkroot := "/hbase"
	client := hbase.NewClient(zkhosts, zkroot)
	//client.SetLogLevel("DEBUG")
	return client
}

//获取结果
func GetResult(client *hbase.Client, table string, rowkey string) (result *hbase.ResultRow, err error) {
	get := hbase.CreateNewGet([]byte(rowkey))
	result, err = client.Get(table, get)
	if (rowkey != result.Row.String()) {
		err = errors.New("没有rowkey")
	}
	return
}
