/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 
 *
 */
package main

import (
	"log"
	"io/ioutil"
	"path/filepath"
	."github.com/hunterhug/go_tool/util"
	"github.com/hunterhug/go_tool/example/testspider"
)


func main() {
	year := 2016
	filepath := filepath.Join(CurDir(),"../","data",IS(year),"1.html")
	log.Println(filepath)
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	returndata:=testspider.DealFile(content)
	log.Printf("%v",returndata)
}