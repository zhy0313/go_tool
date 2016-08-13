/*
 * Created by 一只尼玛 on 2016/8/13.
 * 功能： 
 *
 */
package main

import (
	. "github.com/hunterhug/go_tool/util"
	"log"
	"io/ioutil"
	"os"
	"github.com/hunterhug/go_tool/example/testspider"
	"github.com/hunterhug/go_tool/dbs/mysql"
	"path/filepath"
)

func main() {
	username := "root"
	password := "6833066"
	ip := "localhost"
	dbname := "student"
	fileext:=".html"
	db := mysql.Open(username, password, ip, dbname)
	for i := 1; i < 20; i++ {
		year := 2016
		filepath := filepath.Join(CurDir(), "../", "data", IS(year),IS(i)+fileext)
		log.Println(filepath)
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Printf("%s\n", filepath + "：失败" + err.Error())
			continue
		}
		returndata := testspider.DealFile(content)
		//log.Printf("%v",returndata)
		prestring := "INSERT INTO `info`(`year`,`company`,`name`,`school`,`major`,`grade`,`jobtime`) VALUES('2016',?,?,?,?,?,?)"
		for _, returndata1 := range returndata {
			_, err = mysql.Insert(db, prestring, returndata1)
			if err != nil {
				log.Printf("%v,%s\n", returndata1, filepath + "：失败" + err.Error())
				break
			}else {
				log.Printf("%v\n", returndata1)
			}
		}
		err = os.Rename(filepath, filepath + "lock")
		if err != nil {
			log.Printf("%s\n", filepath + "：失败" + err.Error())
			break
		}

	}
	defer db.Close()
}