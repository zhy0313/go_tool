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
	sql := "INSERT INTO `51job_keyword`(`keyword`,`address`,`kind`) values(?,?,?)"

	num, err := db.Insert(sql, "PHP", "潮州", 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("插入条数%d\n", num)
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
