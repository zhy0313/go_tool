package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

/*　测试之前先填充一下cassandra语句
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
func TestCdb(t *testing.T) {
	//先构造一个字符串数组连接
	host := []string{"192.168.11.74"}
	//指定cassandra keyspace，类似于mysql中的db
	keyword := "example"
	keyword1 := "clicki_v4"
	//连接
	cdb := NewCdb(host, keyword)
	cdb1 := NewCdb(host, keyword1)
	Convey("测试表error", t, func() {
		var unix int
		sql := `SELECT  unix  FROM  info_error WHERE  "id" = ?  and "view_id" = ?   LIMIT 1`
		query := cdb1.Query(sql, 233968, 10)
		iter := query.Iter()
		iter.Scan(&unix)
		err := iter.Close()
		fmt.Printf("unix:%d,%v", unix, err)
	})
	Convey("测试广告表", t, func() {
		var s1 string
		var s2 string
		var s3 string
		sql := `select "campaign_id","spot_id_str","spot_id_str" from info_mz_campaign_spot where "id"=? LIMIT 1`;
		query := cdb1.Query(sql, fmt.Sprintf("%d_%d",1,33))
		iter := query.Iter()
		iter.Scan(&s1, &s2, &s3)
		err := iter.Close()
		fmt.Printf("unix:%s,%s,%s,%v", s1, s2, s3, err)
	})
	Convey("测试cassandra", t, func() {

		//构造插入语句
		insertsql := cdb.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
			"me", gocql.TimeUUID(), "hello wor")

		//执行,查看Ｅxec方法,可知执行后已经关闭
		if err := insertsql.Exec(); err != nil {
			log.Fatal(err)
		}

		//构造查找语句
		querysql := cdb.Query(`SELECT "id", text FROM "tweet" WHERE timeline =?`, "me")

		//执行
		iter := querysql.Iter()

		//定义字节数组
		var id gocql.UUID
		var text string

		//循环取值，需要手工，无法再封装
		if iter.Scan(&id, &text) {
			fmt.Println("\nTweet:", id, text)
		} else {
			fmt.Println("有错误")
		}
		//这个需要关闭
		if err := iter.Close(); err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Println(fmt.Sprintf(`SELECT "%s" FROM "%s" WHERE "%s"=? LIMIT 1`, "2", "RVF", "ID"))
	})
}

