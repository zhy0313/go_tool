# Go_tool
>This is a tool library for Golang.Dont't worry about not understant it!
>All comment writes by English,Ahaha~~ 

>Oh,I think some will be Chinese.

# Usage
```
go get -u -v github.com/hunterhug/go_tool
```

# Include
## Image 图像处理库 **(image deal library)**

```
package image

import (
	"testing"
	"fmt"
)

func TestImage(t *testing.T) {
	err := ThumbnailF2F("../data/image.png", "../data/image100*100.png", 100, 100)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ScaleF2F("../data/image.png", "../data/image200.png", 200)
	if err != nil {
		fmt.Println(err.Error())
	}

	filename, err := RealImageName("../data/image.png")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("real filename"+filename)
	}
}

```

## Spider 爬虫封装库 **(spider library)**

```
package spider

import (
	"testing"
)

func TestSpider(t *testing.T) {
	spiders := Spider{}
	spiders.Method = "get"
	spiders.Wait = 2
	spiders.Url = "http://www.lenggirl.com"

	spiders.SetLogLevel("DEBUg")

	spiders.NewHeader("", "www.baidu.com", "")
	body, err := spiders.Go()
	if err != nil {
		spiders.Log().Error(err)
	} else {
		spiders.Log().Infof("%s", string(body))
	}
	err = TooSortSizes(body, 500)
	spiders.Log().Error(err.Error())
}

```
## Dbs   数据库封装库 **(database library)**

Mysql
```
package mysql

import (
	"testing"
	"fmt"
)

/*

CREATE TABLE `51job_keyword` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `keyword` varchar(255) NOT NULL DEFAULT '',
  `address` varchar(255) NOT NULL DEFAULT '',
  `kind` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `time51` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='关键字表';

*/

func TestMysql(t *testing.T) {

	config := MysqlConfig{
		Username:"root",
		Password:"6833066",
		Ip:"127.0.0.1",
		Port:"3306",
		Dbname:"51job",
	}

	db := New(config)

	db.Open()

	//'1', '教师', '潮州', '0', '2016-05-27 00:00:00', '2016-05-27 00:00:00', '204'
	sql:="INSERT INTO `51job_keyword`(`keyword`,`address`,`kind`) values(?,?,?)"

	num,err:=db.Insert(sql,"PHP","潮州",0)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Printf("插入条数%d\n",num)
	}

	sql = "SELECT * FROM 51job_keyword where address=? and kind=? limit ?;"
	result, err := db.Select(sql, "潮州", 0, 6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for row, v := range result {
			fmt.Printf("%v:%#v\n", row, v)
		}
	}
}

```

## Http 网络库　**(network library)**

## Util 文件/时间等杂项库 **(some small library)**

```
package util

import (
	"testing"
	"fmt"
)

func TestUtil(t *testing.T) {
	i := 2
	if ("2" == IS(i)) {
		fmt.Println("int to string")
	}

	v, err := SI("e2")
	if (err == nil&&v == i) {
		fmt.Println("string to int")
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println(CurDir())

	err = MakeDir("../data")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("dir already exist")
	}

	err = MakeDirByFile("../data/testutil.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("dir already exist")
	}

	err = SaveTofile("../data/testutil.txt", []byte("testutil"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

```

# How to use
>You all can read the test golang file.And I recomment use IDE **pycharm** which python language use,
can also install The Go plugin.

# Author
>一只尼玛

>My website:http://www.lenggirl.com
 
><p>Updating...