/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 
 *
 */
package main

import(
	"github.com/hunterhug/go_tool/dbs/mysql"
	"log"
)
func main() {
	username:="root"
	password:="6833066"
	ip:="localhost"
	dbname:="doubanbook"
	db:=mysql.Open(username,password,ip,dbname)
	defer db.Close()

	prestring:="insert into `book`(`bookno`,`bookurl`,`bookname`,`bookinfo`,`bookstar`) VALUES('d你好','dd',?,?,?)"
	num,err:=mysql.Insert(db,prestring,[]interface{}{"3ddd33","333","0.2"})
	if err!=nil{
		log.Println("dd"+err.Error())
	}
	log.Println(num)
	prestring="SELECT bookname,bookinfo,bookstar FROM book where bookstar = ? limit 10"

	data,err:=mysql.Select(db,prestring,[]interface{}{0.2})
	if err!=nil{
		log.Println(err.Error())
	}
	log.Printf("%v",data)

}
