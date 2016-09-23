package myhbase

import (
	"github.com/hunterhug/go-hbase"
	"log"
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
	"time"
)

func TestHbaseScan(t *testing.T) {
	fs := "2006-01-02 15:04:05"
	Convey("测试scan\n", t, func() {
		client := GetInfoDb()

		//表名
		hbasetable := "dnax:key_info_mz_campaign_spot"

		//列族：列
		hbasecol1 := "clicki-v4:website_name"
		hbasecol2 := "clicki-v4:campaign_name"

		scan := client.Scan(hbasetable)

		//过滤两列，只拿这两列数据
		err := scan.AddString(hbasecol1)

		if err != nil {
			panic(err)
		}
		err = scan.AddString(hbasecol2)
		if err != nil {
			panic(err)
		}

		var t = time.Now().Add(-1 * 720 * time.Hour * 1)
		fmt.Printf("寻找时间在%v之后的记录\n", t.Format(fs))

		//过滤时间戳，只拿时间戳在这段范围的合数据
		//scan.SetTimeRangeFrom(t)

		var v1, v2 string

		//循环处理值
		scan.Map(func(r *hbase.ResultRow) {

			if v, exist := r.Columns[hbasecol1]; exist {
				v1 = v.Value.String()
			}
			if v, exist := r.Columns[hbasecol2]; exist {
				v2 = v.Value.String()
				fmt.Printf("时间:%s\t", hbase.LongToTime(v.Timestamp.Unix()).Format(fs))
			}
			fmt.Printf("row:%s\t  %s:%s \t %s:%s \t", r.Row.String(), hbasecol1, v1, hbasecol2, v2)
			fmt.Printf("addtime:%s\n", hbase.StringToTime(v2).Format(fs))
		})
	})
}
func TestHbaseCommom(t *testing.T) {
	Convey("测试myhbase", t, func() {
		client := GetInfoDb()
		hbasetable := "dnax:info_title"
		rowkey := "test1dd"
		hbasecol := "clicki-v4:test_qualx"
		result, err := GetResult(client, hbasetable, rowkey)
		if err != nil {
			fmt.Printf("在表 [%s] 中找不到 [%s] 的信息,执行hbase错误:%s", hbasetable, rowkey, err.Error())
		}
		if !bytes.Equal(result.Row, []byte(rowkey)) {
			fmt.Printf("在表 [%s] 中找不到 [%s] 的信息", hbasetable, rowkey)
		}
		if v, exist := result.Columns[hbasecol]; exist {
			fmt.Printf("%s", v.Value.String())
		} else {
			fmt.Printf("在表 [%s] 中 主键[%s] 的信息找不到字段[%s]", hbasetable, rowkey, hbasecol)
		}
	})

	Convey("测试hbase", t, func() {
		//scan 'dnax:info_title'
		zkhosts := []string{"192.168.11.73:2181"}
		zkroot := "/hbase"
		client := hbase.NewClient(zkhosts, zkroot)

		table := "dnax:info_title"
		family := "clicki-v4"
		rowkey := "test1d"
		familycol := "test_qual"
		value := "test_val"

		put := hbase.CreateNewPut([]byte(rowkey))
		put.AddStringValue(family, familycol, value)
		res, err := client.Put(table, put)

		if err != nil {
			panic(err)
		}

		if !res {
			panic("No put results")
		}
		log.Println("Completed put")

		get := hbase.CreateNewGet([]byte(rowkey))
		result, err := client.Get(table, get)

		if err != nil {
			panic(err)
		}
		if !bytes.Equal(result.Row, []byte(rowkey)) {
			panic("No row")
		}

		if !bytes.Equal(result.Columns[family + ":" + familycol].Value, []byte(value)) {
			panic("Value doesn't match")
		}

		row := result.Row
		log.Printf("表名:%v", row)
		columns := result.Columns
		for i, c := range columns {
			log.Printf("列族+子列：%v", i)//c.ColumnName
			log.Printf("列族:%v,子列:%v,值:%v", c.Family, c.Qualifier, c.Value)
		}
		log.Println("Completed get")

		//results, err := client.Gets(table, []*hbase.Get{get})
		//
		//if err != nil {
		//	panic(err)
		//}
		//
		//log.Printf("%#v:%v:%v", results, result.Columns, result.Row)
	})
}